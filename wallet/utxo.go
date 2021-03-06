// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wallet

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/BOXFoundation/boxd/core/chain"
	"github.com/BOXFoundation/boxd/core/txlogic"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/rpc/pb"
	"github.com/BOXFoundation/boxd/script"
	"github.com/BOXFoundation/boxd/storage"
	sk "github.com/BOXFoundation/boxd/storage/key"
)

const (
	utxoSelUnitCnt = 256
)

var (
	utxoCacheMtx sync.Mutex
)

type scriptPubKeyFilter func(raw []byte) bool

func filterPayToPubKeyHash(raw []byte) bool {
	return script.NewScriptFromBytes(raw).IsPayToPubKeyHash()
}

func filterToken(raw []byte) bool {
	s := script.NewScriptFromBytes(raw)
	return s.IsTokenTransfer() || s.IsTokenIssue()
}

func filterTokenIssue(raw []byte) bool {
	return script.NewScriptFromBytes(raw).IsTokenIssue()
}

func filterTokenTransfer(raw []byte) bool {
	return script.NewScriptFromBytes(raw).IsTokenTransfer()
}

// BalanceFor returns balance amount of an address using balance index
func BalanceFor(addr string, tid *types.TokenID, db storage.Table) (uint64, error) {
	// check addr
	if _, err := types.NewAddress(addr); err != nil {
		return 0, err
	}
	//
	utxos, err := FetchUtxosOf(addr, tid, 0, db)
	logger.Infof("fetch utxos of %s token %+v got %d utxos", addr, tid, len(utxos))
	if err != nil {
		return 0, err
	}
	var balance uint64
	for _, u := range utxos {
		if u == nil || u.IsSpent {
			logger.Warnf("fetch utxos for %s error, utxo: %+v", addr, u)
			continue
		}
		n, _, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warnf("parse utxo %+v token %+v error: %s", u, tid, err)
			continue
		}
		//if (tid != nil && (tidR == nil || *tid != *tidR)) ||
		//	(tid == nil && tidR != nil) {
		//	logger.Errorf("BalanceFor %s token id %+v got error utxos %+v", u)
		//	continue
		//}
		balance += n
	}
	return balance, nil
}

// FetchUtxosOf fetches utxos from db
// NOTE: if total is 0, fetch all utxos
// NOTE: if tokenID is nil, fetch box utxos
func FetchUtxosOf(
	addr string, tid *types.TokenID, total uint64, db storage.Table,
) ([]*rpcpb.Utxo, error) {

	var utxoKey []byte
	if tid == nil {
		utxoKey = chain.AddrAllUtxoKey(addr)
	} else {
		utxoKey = chain.AddrAllTokenUtxoKey(addr, *tid)
	}
	//
	start := time.Now()
	keys := db.KeysWithPrefix(utxoKey)
	logger.Infof("get utxos keys[%d] for %s amount %d cost %v", len(keys), addr,
		total, time.Since(start))
	// fetch all utxos if total equals to 0
	if total == 0 {
		utxos, err := makeUtxosFromDB(keys, tid, db)
		if err != nil {
			return nil, err
		}
		return utxos, nil
	}
	// fetch moderate utxos by adjustint to total
	utxos, err := fetchModerateUtxos(keys, tid, total, db)
	if err != nil {
		return nil, err
	}
	logger.Infof("fetch utxos for %s amount %d get %d utxos", addr, total, len(utxos))

	return utxos, nil
}

func fetchModerateUtxos(
	keys [][]byte, tid *types.TokenID, total uint64, db storage.Table,
) ([]*rpcpb.Utxo, error) {

	utxoLiveCache.Shrink()
	result := make([]*rpcpb.Utxo, 0)
	remain := total
	for start := 0; start < len(keys) && remain <= total; start += utxoSelUnitCnt {
		// calc start and end keys
		end := start + utxoSelUnitCnt
		if end > len(keys) {
			end = len(keys)
		}
		// fetch utxo from db
		origUtxos, err := makeUtxosFromDB(keys[start:end], tid, db)
		if err != nil {
			return nil, err
		}
		// filter utxos in cache
		utxos := make([]*rpcpb.Utxo, 0, len(origUtxos))
		utxoCacheMtx.Lock()
		for _, u := range origUtxos {
			if utxoLiveCache.Contains(txlogic.ConvPbOutPoint(u.OutPoint)) {
				continue
			}
			utxos = append(utxos, u)
		}
		// select utxos
		selUtxos, amount := selectUtxos(utxos, tid, remain)
		// add utxos to LiveUtxoCache
		for _, u := range utxos {
			utxoLiveCache.Add(txlogic.ConvPbOutPoint(u.OutPoint))
		}
		utxoCacheMtx.Unlock()

		remain -= amount
		result = append(result, selUtxos...)
	}

	return result, nil
}

func makeUtxosFromDB(
	keys [][]byte, tid *types.TokenID, db storage.Table,
) ([]*rpcpb.Utxo, error) {

	ts := time.Now()
	values, err := db.MultiGet(keys...)
	logger.Infof("get utxos values[%d] from db cost %v", len(keys), time.Since(ts))
	if err != nil {
		return nil, err
	}
	// make rpcpb.Utxo array
	utxos := make([]*rpcpb.Utxo, 0, len(values))
	for i, value := range values {
		if value == nil {
			logger.Warnf("utxo not found for key = %s", string(keys[i]))
			continue
		}
		// utxoWrap := new(types.UtxoWrap)
		// if err := utxoWrap.Unmarshal(value); err != nil {
		// 	logger.Warnf("unmarshal error %s, key = %s, body = %v",
		// 		err, string(keys[i]), string(value))
		// 	continue
		// }
		var utxoWrap *types.UtxoWrap
		if utxoWrap, err = chain.DeserializeUtxoWrap(value); err != nil {
			logger.Warnf("Deserialize error %s, key = %s, body = %v",
				err, string(keys[i]), string(value))
			continue
		}
		if utxoWrap == nil {
			logger.Warnf("invalid utxo in db, key: %s, value: %+v", keys[i], utxoWrap)
			continue
		}
		// check utxo type
		spk := utxoWrap.Script()
		var filter scriptPubKeyFilter = filterPayToPubKeyHash
		if tid != nil {
			filter = filterToken
		}
		if !filter(spk) {
			continue
		}
		// make OutPoint
		var op *types.OutPoint
		if tid != nil {
			op, err = parseTokenOutPoint(keys[i])
		} else {
			op, err = parseOutPointFromDbKey(keys[i])
		}
		if err != nil {
			logger.Warn(err)
			continue
		}
		// check utxo token id
		if tid != nil {
			if filterTokenIssue(spk) {
				if *tid != types.TokenID(*op) {
					logger.Warnf("tid: %+v, op: %+v", tid, op)
					continue
				}
			} else if filterTokenTransfer(spk) {
				s := script.NewScriptFromBytes(spk)
				param, err := s.GetTransferParams()
				if err != nil {
					logger.Warn(err)
					continue
				}
				if *tid != types.TokenID(param.TokenID.OutPoint) {
					continue
				}
			} else {
				// other cases, ignore
				continue
			}
		}
		utxos = append(utxos, txlogic.MakePbUtxo(op, utxoWrap))
	}
	return utxos, nil
}

func selectUtxos(
	utxos []*rpcpb.Utxo, tid *types.TokenID, amount uint64,
) ([]*rpcpb.Utxo, uint64) {

	total := uint64(0)
	for _, u := range utxos {
		amount, tidR, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warn(err)
			continue
		}
		if (tid != nil && (tidR == nil || *tid != *tidR)) ||
			(tid == nil && tidR != nil) {
			logger.Errorf("BalanceFor %s token id %+v got error utxos %+v", u)
			continue
		}
		total += amount
	}
	if total <= amount {
		return utxos, total
	}
	// sort
	if tid == nil {
		sort.Sort(sort.Interface(txlogic.SortByUTXOValue(utxos)))
	} else {
		sort.Sort(sort.Interface(txlogic.SortByTokenUTXOValue(utxos)))
	}
	// select
	i, total := 0, uint64(0)
	for k := 0; k < len(utxos) && total < amount; k++ {
		// filter utxos already in cache
		// have check tid and err in the front
		amount, _, _ := txlogic.ParseUtxoAmount(utxos[i])
		total += amount
		i++
	}
	return utxos[:i], total
}

func parseOutPointFromDbKey(key []byte) (*types.OutPoint, error) {
	segs := sk.NewKeyFromBytes(key).List()
	if len(segs) < 4 {
		return nil, fmt.Errorf("invalid address utxo db key %s", string(key))
	}
	return parseOutPointFromKeys(segs[2:4])
}

func parseTokenIDFromDbKey(key []byte) (*types.TokenID, error) {
	segs := sk.NewKeyFromBytes(key).List()
	if len(segs) != 6 {
		return nil, fmt.Errorf("invalid address token utxo db key %s", string(key))
	}
	op, err := parseOutPointFromKeys(segs[2:4])
	return (*types.TokenID)(op), err
}

func parseTokenOutPoint(key []byte) (*types.OutPoint, error) {
	segs := sk.NewKeyFromBytes(key).List()
	if len(segs) != 6 {
		return nil, fmt.Errorf("invalid address token utxo db key %s", string(key))
	}
	return parseOutPointFromKeys(segs[4:6])
}

func parseOutPointFromKeys(segs []string) (*types.OutPoint, error) {
	if len(segs) < 2 {
		return nil, fmt.Errorf("connot parse out point from keys %v", segs)
	}
	hash := new(crypto.HashType)
	if err := hash.SetString(segs[0]); err != nil {
		return nil, err
	}
	index, err := strconv.ParseUint(segs[1], 16, 32)
	if err != nil {
		return nil, err
	}
	return types.NewOutPoint(hash, uint32(index)), nil
}

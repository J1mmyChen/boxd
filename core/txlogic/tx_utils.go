// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package txlogic

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/log"
	"github.com/BOXFoundation/boxd/rpc/pb"
	"github.com/BOXFoundation/boxd/script"
	acc "github.com/BOXFoundation/boxd/wallet/account"
	base58 "github.com/jbenet/go-base58"
)

var logger = log.NewLogger("txlogic") // logger

// NewTokenTag news a TokenTag
func NewTokenTag(name, sym string, decimal uint32, supply uint64) *rpcpb.TokenTag {
	return &rpcpb.TokenTag{
		Name:    name,
		Symbol:  sym,
		Decimal: decimal,
		Supply:  supply,
	}
}

// NewTokenID constructs a token id
func NewTokenID(hash *crypto.HashType, index uint32) *types.TokenID {
	return (*types.TokenID)(types.NewOutPoint(hash, index))
}

// SortByUTXOValue defines a type suited for sort
type SortByUTXOValue []*rpcpb.Utxo

func (x SortByUTXOValue) Len() int           { return len(x) }
func (x SortByUTXOValue) Less(i, j int) bool { return x[i].TxOut.Value < x[j].TxOut.Value }
func (x SortByUTXOValue) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// SortByTokenUTXOValue defines a type suited for sort
type SortByTokenUTXOValue []*rpcpb.Utxo

func (x SortByTokenUTXOValue) Len() int      { return len(x) }
func (x SortByTokenUTXOValue) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x SortByTokenUTXOValue) Less(i, j int) bool {
	vi, err := ParseTokenAmount(x[i].TxOut.GetScriptPubKey())
	if err != nil {
		logger.Warn(err)
	}
	vj, err := ParseTokenAmount(x[j].TxOut.GetScriptPubKey())
	if err != nil {
		logger.Warn(err)
	}
	return vi < vj
}

// ParseUtxoAmount parse amount from utxo and return amount, is token
func ParseUtxoAmount(utxo *rpcpb.Utxo) (uint64, *types.TokenID, error) {
	scp := utxo.TxOut.GetScriptPubKey()
	s := script.NewScriptFromBytes(scp)
	if s.IsPayToPubKeyHash() ||
		s.IsPayToPubKeyHashCLTVScript() ||
		s.IsPayToScriptHash() {
		return utxo.TxOut.GetValue(), nil, nil
	} else if s.IsTokenIssue() {
		tid := (*types.TokenID)(ConvPbOutPoint(utxo.OutPoint))
		amount, err := ParseTokenAmount(scp)
		return amount, tid, err
	} else if s.IsTokenTransfer() {
		param, err := s.GetTransferParams()
		if err != nil {
			return 0, nil, err
		}
		tid := (*types.TokenID)(&param.TokenID.OutPoint)
		return param.Amount, tid, nil
	} else if s.IsSplitAddrScript() {
		return 0, nil, nil
	}
	return 0, nil, errors.New("utxo not recognized")
}

// ParseTokenAmount parse token amount from script pubkey
func ParseTokenAmount(spk []byte) (uint64, error) {
	s := script.NewScriptFromBytes(spk)
	var v uint64
	if s.IsTokenIssue() {
		param, err := s.GetIssueParams()
		if err != nil {
			return 0, err
		}
		v = param.TotalSupply * uint64(math.Pow10(int(param.Decimals)))
	} else if s.IsTokenTransfer() {
		param, err := s.GetTransferParams()
		if err != nil {
			return 0, err
		}
		v = param.Amount
	} else {
		return 0, errors.New("not token script pubkey")
	}
	return v, nil
}

// MakeVout makes txOut
func MakeVout(addr string, amount uint64) *corepb.TxOut {
	var address types.Address
	if strings.HasPrefix(addr, "b2") {
		address, _ = types.NewSplitAddress(addr)
	} else {
		address, _ = types.NewAddress(addr)
	}
	addrPkh, _ := types.NewAddressPubKeyHash(address.Hash())
	addrScript := *script.PayToPubKeyHashScript(addrPkh.Hash())
	return &corepb.TxOut{
		Value:        amount,
		ScriptPubKey: addrScript,
	}
}

// MakeVoutWithSPk makes txOut
func MakeVoutWithSPk(amount uint64, scriptPk []byte) *corepb.TxOut {
	return &corepb.TxOut{
		Value:        amount,
		ScriptPubKey: scriptPk,
	}
}

// MakeVin makes txIn
func MakeVin(utxo *rpcpb.Utxo, seq uint32) *types.TxIn {
	hash := new(crypto.HashType)
	copy(hash[:], utxo.GetOutPoint().Hash)
	return &types.TxIn{
		PrevOutPoint: *types.NewOutPoint(hash, utxo.GetOutPoint().GetIndex()),
		ScriptSig:    []byte{},
		Sequence:     seq,
	}
}

// MakePbVin makes txIn
func MakePbVin(utxo *rpcpb.Utxo, seq uint32) *corepb.TxIn {
	return &corepb.TxIn{
		PrevOutPoint: utxo.OutPoint,
		ScriptSig:    []byte{},
		Sequence:     seq,
	}
}

// NewUtxoWrap makes a UtxoWrap
func NewUtxoWrap(addr string, height uint32, value uint64) *types.UtxoWrap {
	address, _ := types.NewAddress(addr)
	addrPkh, _ := types.NewAddressPubKeyHash(address.Hash())
	addrScript := *script.PayToPubKeyHashScript(addrPkh.Hash())

	return types.NewUtxoWrap(value, addrScript, height)
}

// NewIssueTokenUtxoWrap makes a UtxoWrap
func NewIssueTokenUtxoWrap(
	addr string, tag *rpcpb.TokenTag, height uint32,
) (*types.UtxoWrap, error) {
	vout, err := MakeIssueTokenVout(addr, tag)
	if err != nil {
		return nil, err
	}
	return types.NewUtxoWrap(0, vout.GetScriptPubKey(), height), nil
}

// NewTokenUtxoWrap makes a UtxoWrap
func NewTokenUtxoWrap(
	addr string, tid *types.TokenID, height uint32, value uint64,
) (*types.UtxoWrap, error) {
	vout, err := MakeTokenVout(addr, tid, value)
	if err != nil {
		return nil, err
	}
	return types.NewUtxoWrap(0, vout.GetScriptPubKey(), height), nil
}

// NewPbOutPoint constructs a OutPoint
func NewPbOutPoint(hash *crypto.HashType, index uint32) *corepb.OutPoint {
	return &corepb.OutPoint{
		Hash:  (*hash)[:],
		Index: index,
	}
}

// ConvPbOutPoint constructs a types OutPoint
func ConvPbOutPoint(op *corepb.OutPoint) *types.OutPoint {
	if op == nil {
		return nil
	}
	hash := crypto.HashType{}
	copy(hash[:], op.Hash[:])
	return &types.OutPoint{
		Hash:  hash,
		Index: op.Index,
	}
}

// ConvOutPoint constructs a protobuf OutPoint
func ConvOutPoint(op *types.OutPoint) *corepb.OutPoint {
	return &corepb.OutPoint{
		Hash:  op.Hash[:],
		Index: op.Index,
	}
}

// MakePbUtxo make pb.Utxo from Op and utxo wrap
func MakePbUtxo(op *types.OutPoint, uw *types.UtxoWrap) *rpcpb.Utxo {
	s := script.NewScriptFromBytes(uw.Script())
	value := uw.Value()
	if s.IsTokenIssue() || s.IsTokenTransfer() {
		value = 0
	}
	return &rpcpb.Utxo{
		BlockHeight: uw.Height(),
		IsCoinbase:  uw.IsCoinBase(),
		IsSpent:     uw.IsSpent(),
		OutPoint:    NewPbOutPoint(&op.Hash, op.Index),
		TxOut: &corepb.TxOut{
			Value:        value,
			ScriptPubKey: uw.Script(),
		},
	}
}

// SignTxWithUtxos sign tx with utxo
func SignTxWithUtxos(
	tx *types.Transaction, utxos []*rpcpb.Utxo, acc *acc.Account,
) error {
	for i, utxo := range utxos {
		scriptPkBytes := utxo.GetTxOut().GetScriptPubKey()
		sigHash, err := script.CalcTxHashForSig(scriptPkBytes, tx, i)
		if err != nil {
			return err
		}
		sig, err := acc.Sign(sigHash)
		if err != nil {
			return err
		}
		scriptSig := script.SignatureScript(sig, acc.PublicKey())
		tx.Vin[i].ScriptSig = *scriptSig
	}
	return nil
}

// MakeIssueTokenScript make issue token script for addr with supply and tokent ag
func MakeIssueTokenScript(addr string, tag *rpcpb.TokenTag) ([]byte, error) {
	address, err := types.NewAddress(addr)
	if err != nil {
		return nil, err
	}
	addrPkh, err := types.NewAddressPubKeyHash(address.Hash())
	if err != nil {
		return nil, err
	}
	issueParams := &script.IssueParams{
		Name:        tag.Name,
		Symbol:      tag.Symbol,
		Decimals:    uint8(tag.Decimal),
		TotalSupply: tag.Supply,
	}
	return *script.IssueTokenScript(addrPkh.Hash(), issueParams), nil
}

// MakeIssueTokenVout make issue token vout
func MakeIssueTokenVout(addr string, tag *rpcpb.TokenTag) (*corepb.TxOut, error) {
	spk, err := MakeIssueTokenScript(addr, tag)
	if err != nil {
		return nil, err
	}
	return &corepb.TxOut{Value: 0, ScriptPubKey: spk}, nil
}

// MakeTokenVout make token tx vout
func MakeTokenVout(addr string, tokenID *types.TokenID, amount uint64) (*corepb.TxOut, error) {
	address, err := types.NewAddress(addr)
	if err != nil {
		return nil, err
	}
	addrPkh, err := types.NewAddressPubKeyHash(address.Hash())
	if err != nil {
		return nil, err
	}
	transferParams := &script.TransferParams{}
	transferParams.Hash = tokenID.Hash
	transferParams.Index = tokenID.Index
	transferParams.Amount = amount
	addrScript := *script.TransferTokenScript(addrPkh.Hash(), transferParams)
	return &corepb.TxOut{Value: 0, ScriptPubKey: addrScript}, nil
}

// MakeSplitAddrVout make split addr vout
func MakeSplitAddrVout(addrs []string, weights []uint64) *corepb.TxOut {
	return &corepb.TxOut{
		Value:        0,
		ScriptPubKey: MakeSplitAddrPubkey(addrs, weights),
	}
}

// MakeSplitAddrPubkey make split addr
func MakeSplitAddrPubkey(addrs []string, weights []uint64) []byte {
	addresses := make([]types.Address, len(addrs))
	for i, addr := range addrs {
		addresses[i], _ = types.NewAddress(addr)
	}
	return *script.SplitAddrScript(addresses, weights)
}

// MakeSplitAddr make split addr
func MakeSplitAddr(addrs []string, weights []uint64) (string, error) {
	pk := MakeSplitAddrPubkey(addrs, weights)
	splitAddrScriptStr := script.NewScriptFromBytes(pk).Disasm()
	s := strings.Split(splitAddrScriptStr, " ")
	pubKeyHash, err := hex.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	addr, err := types.NewSplitAddressFromHash(pubKeyHash)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

// IsCoinBaseTxIn check whether tx in is coin base tx in
func IsCoinBaseTxIn(txIn *types.TxIn) bool {
	return ((txIn.PrevOutPoint.Index == math.MaxUint32) &&
		(txIn.PrevOutPoint.Hash == crypto.HashType{}))
}

// NewCoinBaseTxIn new a coinbase tx in
func NewCoinBaseTxIn() *types.TxIn {
	return &types.TxIn{
		PrevOutPoint: types.OutPoint{Index: math.MaxUint32},
	}
}

// EncodeOutPoint encode token to string
func EncodeOutPoint(op *corepb.OutPoint) string {
	buf := make([]byte, len(op.Hash))
	copy(buf, op.Hash[:])
	// reverse bytes
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	// append separator ':'
	buf = append(buf, ':')
	// put index
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, op.Index)
	buf = append(buf, b...)

	return base58.Encode(buf)
}

// DecodeOutPoint string token id to TokenID
func DecodeOutPoint(id string) (*corepb.OutPoint, error) {
	buf := base58.Decode(id)
	if len(buf) != crypto.HashSize+5 {
		return nil, fmt.Errorf("decode tokenID error, length(%d) mismatch, data: %s",
			crypto.HashSize+5, id)
	}
	if buf[crypto.HashSize] != ':' {
		return nil, fmt.Errorf("token id delimiter want ':', got: %c, data: %s",
			buf[crypto.HashSize], id)
	}
	for i, j := 0, crypto.HashSize-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	index := binary.LittleEndian.Uint32(buf[crypto.HashSize+1:])
	hash := new(crypto.HashType)
	hash.SetBytes(buf[:crypto.HashSize])
	return NewPbOutPoint(hash, index), nil
}

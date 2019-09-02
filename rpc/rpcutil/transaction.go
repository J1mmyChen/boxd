// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rpcutil

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/BOXFoundation/boxd/boxd/service"
	"github.com/BOXFoundation/boxd/core"
	corepb "github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/core/txlogic"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	rpcpb "github.com/BOXFoundation/boxd/rpc/pb"
	acc "github.com/BOXFoundation/boxd/wallet/account"
	"google.golang.org/grpc"
)

const (
	connTimeout = 30
	maxDecimal  = 8
)

// GetBalance returns total amount of an address
func GetBalance(conn *grpc.ClientConn, addresses []string) ([]uint64, error) {
	c := rpcpb.NewTransactionCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)
	defer cancel()
	resp, err := c.GetBalance(ctx, &rpcpb.GetBalanceReq{Addrs: addresses})
	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.GetBalances(), nil
}

// GetTokenBalance returns total amount of an address with specified token id
func GetTokenBalance(
	conn *grpc.ClientConn, addresses []string, tokenHash string, tokenIndex uint32,
) ([]uint64, error) {

	c := rpcpb.NewTransactionCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)
	defer cancel()
	resp, err := c.GetTokenBalance(ctx, &rpcpb.GetTokenBalanceReq{
		Addrs: addresses, TokenHash: tokenHash, TokenIndex: tokenIndex})
	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.GetBalances(), nil
}

func newFetchUtxosReq(addr string, amount uint64) *rpcpb.FetchUtxosReq {
	return &rpcpb.FetchUtxosReq{Addr: addr, Amount: amount}
}

// FetchUtxos fetch utxos from chain
func FetchUtxos(
	conn *grpc.ClientConn, addr string, amount uint64, tHashStr string, tIdx uint32,
) ([]*rpcpb.Utxo, error) {

	c := rpcpb.NewTransactionCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)
	defer cancel()
	req := newFetchUtxosReq(addr, amount)
	if tHashStr != "" {
		hash := new(crypto.HashType)
		if err := hash.SetString(tHashStr); err != nil {
			return nil, err
		}
		req.TokenHash = tHashStr
		req.TokenIndex = tIdx
	}
	resp, err := c.FetchUtxos(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.GetUtxos(), nil
}

// NewIssueTokenTx new a issue token transaction
func NewIssueTokenTx(
	acc *acc.Account, to *types.AddressHash, tag *rpcpb.TokenTag, conn *grpc.ClientConn,
) (*types.Transaction, *types.TokenID, *rpcpb.Utxo, error) {

	// fetch utxos for fee
	utxos, err := fetchUtxos(conn, acc.Addr(), core.TransferFee, "", 0)
	if err != nil {
		return nil, nil, nil, err
	}
	inputAmt := uint64(0)
	for _, u := range utxos {
		inputAmt += u.GetTxOut().GetValue()
	}
	//
	tx, tid, change, err := txlogic.NewTokenIssueTxWithUtxos(acc, to, tag,
		inputAmt-core.TransferFee, utxos...)
	if err != nil {
		logger.Warnf("new issue token tx with utxos from %s to %s tag %+v "+
			"supply %d change %d with utxos: %+v error: %s", acc.Addr(), to, tag,
			tag.Supply, inputAmt-core.TransferFee, utxos, err)
		return nil, nil, nil, err
	}
	return tx, tid, change, nil
}

// NewContractDeployTx new a deploy contract transaction
func NewContractDeployTx(
	acc *acc.Account, gasLimit, nonce uint64, code []byte, conn *grpc.ClientConn,
) (*types.Transaction, string, error) {
	// fetch utxos for gas
	utxos, err := fetchUtxos(conn, acc.Addr(), core.FixedGasPrice*gasLimit, "", 0)
	if err != nil {
		return nil, "", err
	}
	inputAmt := uint64(0)
	for _, u := range utxos {
		inputAmt += u.GetTxOut().GetValue()
	}
	//
	tx, err := txlogic.MakeUnsignedContractDeployTx(acc.AddressHash(), 0,
		inputAmt-core.FixedGasPrice*gasLimit, gasLimit, nonce, code, utxos...)
	if err != nil {
		return nil, "", err
	}
	// sign vin
	if err = txlogic.SignTxWithUtxos(tx, utxos, acc); err != nil {
		return nil, "", err
	}
	// nonce

	senderAddr, err := types.NewAddress(acc.Addr())
	if err != nil {
		return nil, "", err
	}
	contractAddr, _ := types.MakeContractAddress(senderAddr, nonce)
	return tx, contractAddr.String(), nil
}

// NewContractCallTx new a call contract transaction
func NewContractCallTx(
	acc *acc.Account, to *types.AddressHash, gasLimit, nonce uint64, code []byte,
	conn *grpc.ClientConn,
) (*types.Transaction, error) {
	// fetch utxos for gas
	utxos, err := fetchUtxos(conn, acc.Addr(), core.FixedGasPrice*gasLimit, "", 0)
	if err != nil {
		return nil, err
	}
	inputAmt := uint64(0)
	for _, u := range utxos {
		inputAmt += u.GetTxOut().GetValue()
	}
	//
	tx, err := txlogic.MakeUnsignedContractCallTx(acc.AddressHash(), to, 0,
		inputAmt-core.FixedGasPrice*gasLimit, gasLimit, nonce, code, utxos...)
	if err != nil {
		return nil, err
	}
	// sign vin
	if err = txlogic.SignTxWithUtxos(tx, utxos, acc); err != nil {
		return nil, err
	}
	return tx, nil
}

// SendTransaction sends an signed transaction to node server through grpc connection
func SendTransaction(conn *grpc.ClientConn, tx *types.Transaction) (string, error) {
	txProtoMsg, err := tx.ToProtoMessage()
	if err != nil {
		return "", err
	}
	txPb, ok := txProtoMsg.(*corepb.Transaction)
	if !ok {
		return "", fmt.Errorf("can't convert transaction into protobuf")
	}
	txReq := &rpcpb.SendTransactionReq{Tx: txPb}

	c := rpcpb.NewTransactionCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)
	defer cancel()

	resp, err := c.SendTransaction(ctx, txReq)
	if err != nil {
		return "", err
	} else if resp.GetCode() != 0 {
		return "", errors.New(resp.GetMessage())
	}
	return resp.Hash, nil
}

//CreateRawTransaction create a tx without signature,it returns a tx and utxo
func CreateRawTransaction(
	from *types.AddressHash, to []*types.AddressHash, txhash []crypto.HashType,
	vout []uint32, amounts []uint64, height uint32,
) (*types.Transaction, error) {
	total := uint64(0)
	for _, a := range amounts {
		total += a
	}

	utxos := make([]*rpcpb.Utxo, 0)
	for i := 0; i < len(txhash); i++ {
		op := types.NewOutPoint(&txhash[i], vout[i])
		uw := txlogic.NewUtxoWrap(from, height, total)
		utxo := txlogic.MakePbUtxo(op, uw)
		utxos = append(utxos, utxo)
	}
	changeAmt, _, overflowed := calcChangeAmount(amounts, 0, utxos...)
	if overflowed {
		return nil, txlogic.ErrInsufficientBalance
	}
	tx, err := txlogic.MakeUnsignedTx(from, to, amounts, changeAmt, utxos...)
	return tx, err
}

// GetRawTransaction get the transaction info of given hash
func GetRawTransaction(conn *grpc.ClientConn, hash []byte) (*types.Transaction, error) {
	c := rpcpb.NewTransactionCommandClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout*time.Second)
	defer cancel()
	logger.Debugf("Get transaction of hash: %x", hash)

	r, err := c.GetRawTransaction(ctx, &rpcpb.GetRawTransactionRequest{Hash: hash})
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{}
	err = tx.FromProtoMessage(r.Tx)
	return tx, err
}

// NewTx new a tx and return change utxo
func NewTx(
	fromAcc *acc.Account, toAddrs []*types.AddressHash, amounts []uint64,
	conn *grpc.ClientConn,
) (tx *types.Transaction, change *rpcpb.Utxo, err error) {
	// calc amount
	amount := uint64(0)
	for _, a := range amounts {
		amount += a
	}
	// get utxos
	utxos, err := fetchUtxos(conn, fromAcc.Addr(), amount+core.TransferFee, "", 0)
	if err != nil {
		err = fmt.Errorf("fetchUtxos error for %s amount %d: %s",
			fromAcc.Addr(), amount+core.TransferFee, err)
		return
	}
	// NOTE: for test only
	//checkDuplicateUtxos(utxos)
	// calc change amount
	total := uint64(0)
	for _, u := range utxos {
		total += u.GetTxOut().GetValue()
	}
	changeAmt := total - amount - core.TransferFee
	if changeAmt >= total {
		err = fmt.Errorf("invalid arguments, addr %s utxo total=%d, amount=%d, "+
			"changeAmt=%d", fromAcc.Addr(), total, amount, changeAmt)
		return
	}
	//
	tx, change, err = txlogic.NewTxWithUtxos(fromAcc, toAddrs, utxos, amounts, changeAmt)
	return
}

// NewTxs construct some transactions
func NewTxs(
	fromAcc *acc.Account, toAddr *types.AddressHash, count int, conn *grpc.ClientConn,
) (txss [][]*types.Transaction, transfer, totalFee uint64, num int, err error) {

	// get utxoes
	utxos, err := fetchUtxos(conn, fromAcc.Addr(), 0, "", 0)
	if err != nil {
		return
	}
	if len(utxos) == 0 {
		err = fmt.Errorf("no utxos")
		return
	}
	// NOTE: for test only
	//checkDuplicateUtxos(utxos)
	// gen txs
	txss = make([][]*types.Transaction, 0)
	n := (count + len(utxos) - 1) / len(utxos)

	for _, u := range utxos {
		change := u
		value := change.GetTxOut().GetValue()
		aveAmt := value / uint64(n)
		if aveAmt == 0 {
			continue
		}
		changeAmt := value
		txs := make([]*types.Transaction, 0)
		for j := n; num < count && j > 0; j-- {
			amount := aveAmt - core.TransferFee
			changeAmt = changeAmt - aveAmt
			tx := new(types.Transaction)
			tx, change, err = txlogic.NewTxWithUtxos(fromAcc, []*types.AddressHash{toAddr},
				[]*rpcpb.Utxo{change}, []uint64{amount}, changeAmt)
			if err != nil {
				logger.Warn(err)
				continue
			}
			txs = append(txs, tx)
			transfer += amount
			totalFee += core.TransferFee
			num++
			if change == nil {
				break
			}
		}
		txss = append(txss, txs)
	}
	return txss, transfer, totalFee, num, nil
}

func fetchUtxos(
	conn *grpc.ClientConn, addr string, amount uint64, tHashStr string, tIdx uint32,
) (utxos []*rpcpb.Utxo, err error) {
	for t := 0; t < 30; t++ {
		utxos, err = FetchUtxos(conn, addr, amount, tHashStr, tIdx)
		if len(utxos) == 0 {
			err = fmt.Errorf("fetch no utxo for %s amount %d", addr, amount)
		}
		if err == nil {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}
	if err != nil {
		logger.Warn(err)
	}
	return utxos, err
}

// GetGRPCConn returns a conn with peer addr
func GetGRPCConn(peerAddr string) (*grpc.ClientConn, error) {
	return grpc.Dial(peerAddr, grpc.WithInsecure())
}

// NewTokenTx new a token tx
func NewTokenTx(
	acc *acc.Account, toAddrs []*types.AddressHash, amounts []uint64, tHashStr string,
	tIdx uint32, conn *grpc.ClientConn,
) (*types.Transaction, *rpcpb.Utxo, *rpcpb.Utxo, error) {

	amount := core.TransferFee
	amountT := uint64(0)
	for _, a := range amounts {
		amountT += a
	}
	boxUtxos, err := fetchUtxos(conn, acc.Addr(), amount, "", 0)
	if err != nil {
		return nil, nil, nil, err
	}
	tokenUtxos, err := fetchUtxos(conn, acc.Addr(), amountT, tHashStr, tIdx)
	if err != nil {
		return nil, nil, nil, err
	}
	utxos := append(boxUtxos, tokenUtxos...)
	tHash := new(crypto.HashType)
	if err := tHash.SetString(tHashStr); err != nil {
		return nil, nil, nil, err
	}
	tid := (*types.TokenID)(types.NewOutPoint(tHash, tIdx))
	return txlogic.NewTokenTransferTxWithUtxos(acc, toAddrs, amounts, tid, 0, utxos...)
}

// NewTokenTxs new a token tx
func NewTokenTxs(
	acc *acc.Account, toAddr *types.AddressHash, amountT uint64, count int, tHashStr string,
	tIdx uint32, conn *grpc.ClientConn,
) ([]*types.Transaction, error) {
	// get utxos
	amount := core.TransferFee * uint64(count)
	boxUtxos, err := fetchUtxos(conn, acc.Addr(), amount, "", 0)
	if err != nil {
		return nil, err
	}
	tokenUtxos, err := fetchUtxos(conn, acc.Addr(), amountT, tHashStr, tIdx)
	if err != nil {
		return nil, err
	}
	utxos := append(boxUtxos, tokenUtxos...)
	// tid
	tHash := new(crypto.HashType)
	if err := tHash.SetString(tHashStr); err != nil {
		return nil, err
	}
	tid := (*types.TokenID)(types.NewOutPoint(tHash, tIdx))
	//
	var txs []*types.Transaction
	unitT := amountT / uint64(count)
	changeAmt := amount
	for i := 0; i < count; i++ {
		if i == count-1 {
			unitT = amountT - unitT*uint64(i)
		}
		changeAmt -= core.TransferFee
		tx, change, changeT, err := txlogic.NewTokenTransferTxWithUtxos(acc,
			[]*types.AddressHash{toAddr}, []uint64{unitT}, tid, changeAmt, utxos...)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
		utxos = []*rpcpb.Utxo{change, changeT}
		if change == nil {
			break
		}
	}
	return txs, nil
}

// NewERC20TransferFromContractTxs new a contract transferFrom tx
func NewERC20TransferFromContractTxs(
	acc *acc.Account, contractAddr *types.AddressHash, count int, gasLimit, startNonce uint64,
	code []byte, conn *grpc.ClientConn,
) ([]*types.Transaction, error) {
	// get utxos
	amount := core.FixedGasPrice * gasLimit * uint64(count)
	utxos, err := fetchUtxos(conn, acc.Addr(), amount, "", 0)
	if err != nil {
		return nil, err
	}
	var boxAmt uint64
	for _, u := range utxos {
		boxAmt += u.TxOut.Value
	}
	//
	var txs []*types.Transaction
	changeAmt, nonce := boxAmt, startNonce
	for i := 0; i < count; i++ {
		changeAmt -= core.FixedGasPrice * gasLimit
		tx, change, err := txlogic.NewContractTxWithUtxos(acc, contractAddr, 0,
			changeAmt, gasLimit, nonce, code, utxos...)
		if err != nil {
			return nil, err
		}
		nonce++
		txs = append(txs, tx)
		utxos = []*rpcpb.Utxo{change}
		if change == nil {
			break
		}
	}
	return txs, nil
}

// NewSplitAddrTx new split address tx
func NewSplitAddrTx(
	acc *acc.Account, addrs []*types.AddressHash, weights []uint32, conn *grpc.ClientConn,
) (tx *types.Transaction, change *rpcpb.Utxo, err error) {
	// get utxos
	utxos, err := fetchUtxos(conn, acc.Addr(), core.TransferFee, "", 0)
	if err != nil {
		return
	}
	return txlogic.NewSplitAddrTxWithUtxos(acc, addrs, weights, utxos)
}

// MakeUnsignedTx make a tx without signature
func MakeUnsignedTx(
	wa service.WalletAgent, from *types.AddressHash, to []*types.AddressHash,
	amounts []uint64, gasUsed uint64,
) (*types.Transaction, []*rpcpb.Utxo, error) {
	total := gasUsed
	for _, a := range amounts {
		total += a
	}
	utxos, err := wa.Utxos(from, nil, total)
	if err != nil {
		return nil, nil, err
	}
	changeAmt, _, overflowed := calcChangeAmount(amounts, gasUsed, utxos...)
	if overflowed {
		return nil, nil, txlogic.ErrInsufficientBalance
	}
	tx, err := txlogic.MakeUnsignedTx(from, to, amounts, changeAmt, utxos...)
	return tx, utxos, err
}

// MakeUnsignedCombineTx make a combine tx without signature
func MakeUnsignedCombineTx(
	wa service.WalletAgent, from *types.AddressHash, gasUsed uint64,
) (*types.Transaction, []*rpcpb.Utxo, error) {

	utxos, err := wa.Utxos(from, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	utxosAmt := uint64(0)
	for _, u := range utxos {
		amount, _, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warn(err)
			continue
		}
		utxosAmt += amount
	}
	tx, err := txlogic.MakeUnsignedTx(from, []*types.AddressHash{from}, []uint64{utxosAmt - gasUsed},
		0, utxos...)
	return tx, utxos, err
}

// MakeUnsignedCombineTokenTx make a combine tx without signature
func MakeUnsignedCombineTokenTx(
	wa service.WalletAgent, from *types.AddressHash, tid *types.TokenID, gasUsed uint64,
) (*types.Transaction, []*rpcpb.Utxo, error) {

	// fetch box utxos
	utxos, err := wa.Utxos(from, nil, gasUsed)
	if err != nil {
		return nil, nil, err
	}
	utxosAmt := uint64(0)
	for _, u := range utxos {
		amount, _, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warn(err)
			return nil, nil, err
		}
		utxosAmt += amount
	}
	totalUtxos := utxos
	// fetch token utxos
	utxos, err = wa.Utxos(from, tid, 0)
	if err != nil {
		return nil, nil, err
	}
	utxosAmtT := uint64(0)
	for _, u := range utxos {
		amount, _, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warn(err)
			return nil, nil, err
		}
		utxosAmtT += amount
	}
	totalUtxos = append(totalUtxos, utxos...)
	// make tx
	tx, _, err := txlogic.MakeUnsignedTokenTransferTx(from, []*types.AddressHash{from},
		[]uint64{utxosAmtT}, tid, utxosAmt-gasUsed, totalUtxos...)
	return tx, utxos, err
}

//MakeUnsignedContractDeployTx make a tx without a signature
func MakeUnsignedContractDeployTx(
	wa service.WalletAgent, from *types.AddressHash, amount, gasLimit, nonce uint64,
	byteCode []byte,
) (*types.Transaction, []*rpcpb.Utxo, error) {
	gasUsed := gasLimit * core.FixedGasPrice
	total := gasUsed + amount
	utxos, err := wa.Utxos(from, nil, total)
	if err != nil {
		return nil, nil, err
	}
	amounts := []uint64{amount}
	changeAmt, _, overflowed := calcChangeAmount(amounts, gasUsed, utxos...)
	if overflowed {
		return nil, nil, txlogic.ErrInsufficientBalance
	}
	if err != nil {
		return nil, nil, err
	}
	tx, err := txlogic.MakeUnsignedContractDeployTx(from, amount, changeAmt,
		gasLimit, nonce, byteCode, utxos...)
	return tx, utxos, err
}

//MakeUnsignedContractCallTx call a contract tx without a signature
func MakeUnsignedContractCallTx(
	wa service.WalletAgent, from *types.AddressHash, amount, gasLimit, nonce uint64,
	contractAddr *types.AddressHash, byteCode []byte,
) (*types.Transaction, []*rpcpb.Utxo, error) {
	gasUsed := gasLimit * core.FixedGasPrice
	total := gasUsed + amount
	utxos, err := wa.Utxos(from, nil, total)
	if err != nil {
		return nil, nil, err
	}
	amounts := []uint64{amount}
	changeAmt, _, overflowed := calcChangeAmount(amounts, gasUsed, utxos...)
	if overflowed {
		return nil, nil, txlogic.ErrInsufficientBalance
	}
	if err != nil {
		return nil, nil, err
	}
	tx, err := txlogic.MakeUnsignedContractCallTx(from, contractAddr, amount,
		changeAmt, gasLimit, nonce, byteCode, utxos...)
	return tx, utxos, err
}

// MakeUnsignedSplitAddrTx news tx to make split addr without signature
// it returns a tx, split addr, a change
func MakeUnsignedSplitAddrTx(
	wa service.WalletAgent, from *types.AddressHash, addrs []*types.AddressHash,
	weights []uint32, gasUsed uint64,
) (*types.Transaction, []*rpcpb.Utxo, error) {
	utxos, err := wa.Utxos(from, nil, gasUsed)
	if err != nil {
		return nil, nil, err
	}
	changeAmt, _, overflowed := calcChangeAmount(nil, gasUsed, utxos...)
	if overflowed {
		return nil, nil, txlogic.ErrInsufficientBalance
	}
	tx, err := txlogic.MakeUnsignedSplitAddrTx(from, addrs, weights,
		changeAmt, utxos...)
	return tx, utxos, err
}

// MakeUnsignedTokenIssueTx news tx to issue token without signature
func MakeUnsignedTokenIssueTx(
	wa service.WalletAgent, issuer, owner *types.AddressHash, tag *rpcpb.TokenTag,
	gasUsed uint64,
) (*types.Transaction, uint32, []*rpcpb.Utxo, error) {
	if tag.GetDecimal() > maxDecimal {
		return nil, 0, nil, fmt.Errorf("Decimal[%d] is bigger than max Decimal[%d]", tag.GetDecimal(), maxDecimal)
	}
	if tag.GetSupply() > math.MaxUint64/uint64(math.Pow10(int(tag.GetDecimal()))) {
		return nil, 0, nil, fmt.Errorf("supply is too bigger")
	}
	utxos, err := wa.Utxos(issuer, nil, gasUsed)
	if err != nil {
		return nil, 0, nil, err
	}
	changeAmt, _, overflowed := calcChangeAmount(nil, gasUsed, utxos...)
	if overflowed {
		return nil, 0, nil, txlogic.ErrInsufficientBalance
	}
	tx, issueOutIndex, err := txlogic.MakeUnsignedTokenIssueTx(issuer, owner, tag,
		changeAmt, utxos...)
	return tx, issueOutIndex, utxos, err
}

// MakeUnsignedTokenTransferTx news tx to transfer token without signature
func MakeUnsignedTokenTransferTx(
	wa service.WalletAgent, from *types.AddressHash, to []*types.AddressHash,
	amounts []uint64, tid *types.TokenID, gasUsed uint64,
) (*types.Transaction, []*rpcpb.Utxo, error) {
	utxos, err := wa.Utxos(from, nil, gasUsed)
	if err != nil {
		return nil, nil, err
	}
	tAmt := uint64(0)
	for _, a := range amounts {
		tAmt += a
	}
	tokenUtxos, err := wa.Utxos(from, tid, tAmt)
	if err != nil {
		return nil, nil, err
	}
	changeAmt, _, overflowed := calcChangeAmount(nil, gasUsed, utxos...)
	if overflowed {
		return nil, nil, txlogic.ErrInsufficientBalance
	}
	mixUtxos := append(utxos, tokenUtxos...)
	tx, _, err := txlogic.MakeUnsignedTokenTransferTx(from, to, amounts, tid,
		changeAmt, mixUtxos...)

	return tx, mixUtxos, err
}

func calcChangeAmount(
	amounts []uint64, gasUsed uint64, utxos ...*rpcpb.Utxo,
) (changeAmt uint64, utxosAmt uint64, overflowed bool) {
	total := gasUsed
	for _, a := range amounts {
		total += a
	}
	for _, u := range utxos {
		amount, tid, err := txlogic.ParseUtxoAmount(u)
		if err != nil || tid != nil {
			continue
		}
		utxosAmt += amount
	}
	changeAmt = utxosAmt - total
	return changeAmt, utxosAmt, changeAmt > utxosAmt
}

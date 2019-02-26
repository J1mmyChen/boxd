// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rpc

import (
	"context"
	"errors"
	"reflect"

	"github.com/BOXFoundation/boxd/core"
	"github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/core/txlogic"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/rpc/pb"
	"github.com/BOXFoundation/boxd/rpc/rpcutil"
)

func registerTransaction(s *Server) {
	rpcpb.RegisterTransactionCommandServer(s.server, &txServer{server: s})
}

func init() {
	RegisterServiceWithGatewayHandler(
		"tx",
		registerTransaction,
		rpcpb.RegisterTransactionCommandHandlerFromEndpoint,
	)
}

type txServer struct {
	server GRPCServer
}

func newGetBalanceResp(code int32, msg string, balances ...uint64) *rpcpb.GetBalanceResp {
	return &rpcpb.GetBalanceResp{
		Code:     code,
		Message:  msg,
		Balances: balances,
	}
}

func (s *txServer) GetBalance(
	ctx context.Context, req *rpcpb.GetBalanceReq,
) (*rpcpb.GetBalanceResp, error) {
	logger.Infof("get balance req: %+v", req)
	walletAgent := s.server.GetWalletAgent()
	if walletAgent == nil || reflect.ValueOf(walletAgent).IsNil() {
		logger.Warn("get balance error ", ErrAPINotSupported)
		return newGetBalanceResp(-1, ErrAPINotSupported.Error()), ErrAPINotSupported
	}
	balances := make([]uint64, len(req.GetAddrs()))
	for i, addr := range req.Addrs {
		amount, err := walletAgent.Balance(addr, nil)
		if err != nil {
			logger.Warnf("get balance for %s error %s", addr, err)
			return newGetBalanceResp(-1, err.Error()), nil
		}
		balances[i] = amount
	}
	logger.Infof("get balance for %v result: %v", req.Addrs, balances)
	return newGetBalanceResp(0, "ok", balances...), nil
}

func (s *txServer) GetTokenBalance(
	ctx context.Context, req *rpcpb.GetTokenBalanceReq,
) (*rpcpb.GetBalanceResp, error) {
	logger.Infof("get token balance req: %+v", req)
	walletAgent := s.server.GetWalletAgent()
	if walletAgent == nil {
		logger.Warn("get token balance error ", ErrAPINotSupported)
		return newGetBalanceResp(-1, ErrAPINotSupported.Error()), ErrAPINotSupported
	}
	balances := make([]uint64, len(req.GetAddrs()))
	tid := (*txlogic.TokenID)(txlogic.ConvPbOutPoint(req.TokenID))
	for i, addr := range req.Addrs {
		amount, err := walletAgent.Balance(addr, tid)
		if err != nil {
			logger.Warnf("get token balance for %s error %s", addr, err)
			return newGetBalanceResp(-1, err.Error()), nil
		}
		balances[i] = amount
	}
	logger.Infof("get balance for %v result: %v", req.Addrs, balances)
	return newGetBalanceResp(0, "ok", balances...), nil
}

func newFetchUtxosResp(code int32, msg string, utxos ...*rpcpb.Utxo) *rpcpb.FetchUtxosResp {
	return &rpcpb.FetchUtxosResp{
		Code:    code,
		Message: msg,
		Utxos:   utxos,
	}
}

// FetchUtxos fetch utxos from chain
func (s *txServer) FetchUtxos(
	ctx context.Context, req *rpcpb.FetchUtxosReq,
) (*rpcpb.FetchUtxosResp, error) {

	logger.Infof("fetch utxos req: %+v", req)
	walletAgent := s.server.GetWalletAgent()
	if walletAgent == nil || reflect.ValueOf(walletAgent).IsNil() {
		logger.Warn("fetch utxos error ", ErrAPINotSupported)
		return newFetchUtxosResp(-1, ErrAPINotSupported.Error()), ErrAPINotSupported
	}
	var tid *txlogic.TokenID
	if req.GetTokenID() != nil {
		tid = (*txlogic.TokenID)(txlogic.ConvPbOutPoint(req.GetTokenID()))
	}
	utxos, err := walletAgent.Utxos(req.GetAddr(), tid, req.GetAmount())
	if err != nil {
		logger.Warnf("fetch utxos for %+v error %s", req, err)
		return newFetchUtxosResp(-1, err.Error()), nil
	}
	total := uint64(0)
	for _, u := range utxos {
		amount, _, err := txlogic.ParseUtxoAmount(u)
		if err != nil {
			logger.Warnf("fetch utxos for %+v error %s", req, err)
			continue
		}
		total += amount
	}
	logger.Infof("fetch utxos for %+v return %d utxos total %d",
		req, len(utxos), total)
	return newFetchUtxosResp(0, "ok", utxos...), nil
}

func (s *txServer) GetTransactionPool(ctx context.Context, req *rpcpb.GetTransactionPoolRequest) (*rpcpb.GetTransactionsResponse, error) {
	txs, _ := s.server.GetTxHandler().GetTransactionsInPool()
	respTxs := []*corepb.Transaction{}
	for _, tx := range txs {
		respTx, err := tx.ToProtoMessage()
		if err != nil {
			return &rpcpb.GetTransactionsResponse{}, err
		}

		respTxs = append(respTxs, respTx.(*corepb.Transaction))
	}
	return &rpcpb.GetTransactionsResponse{Txs: respTxs}, nil
}

func (s *txServer) GetFeePrice(ctx context.Context, req *rpcpb.GetFeePriceRequest) (*rpcpb.GetFeePriceResponse, error) {
	return &rpcpb.GetFeePriceResponse{BoxPerByte: 1}, nil
}

func newSendTransactionResp(code int32, msg, hash string) *rpcpb.SendTransactionResp {
	return &rpcpb.SendTransactionResp{
		Code:    code,
		Message: msg,
		Hash:    hash,
	}
}

func (s *txServer) SendTransaction(
	ctx context.Context, req *rpcpb.SendTransactionReq,
) (resp *rpcpb.SendTransactionResp, err error) {

	logger.Infof("send tx req >> %+v", req)
	defer func() {
		if resp.Code != 0 {
			logger.Warnf("send tx %+v error: %s", req, resp.Message)
		} else {
			logger.Infof("send tx: %+v succeeded, response: %+v", resp)
		}
	}()

	tx := &types.Transaction{}
	if err := tx.FromProtoMessage(req.Tx); err != nil {
		return newSendTransactionResp(-1, err.Error(), ""), nil
	}

	txpool := s.server.GetTxHandler()
	if err := txpool.ProcessTx(tx, core.BroadcastMode); err != nil {
		return newSendTransactionResp(-1, err.Error(), ""), nil
	}
	hash, _ := tx.TxHash()
	return newSendTransactionResp(0, "success", hash.String()), nil
}

func (s *txServer) GetRawTransaction(
	ctx context.Context, req *rpcpb.GetRawTransactionRequest,
) (*rpcpb.GetRawTransactionResponse, error) {
	hash := crypto.HashType{}
	if err := hash.SetBytes(req.Hash); err != nil {
		return &rpcpb.GetRawTransactionResponse{}, err
	}
	tx, err := s.server.GetChainReader().LoadTxByHash(hash)
	if err != nil {
		logger.Debug(err)
		return &rpcpb.GetRawTransactionResponse{}, err
	}
	rpcTx, err := tx.ToProtoMessage()
	return &rpcpb.GetRawTransactionResponse{Tx: rpcTx.(*corepb.Transaction)}, err
}

func generateUtxoMessage(outPoint *types.OutPoint, entry *types.UtxoWrap) *rpcpb.Utxo {
	return &rpcpb.Utxo{
		BlockHeight: entry.Height(),
		IsCoinbase:  entry.IsCoinBase(),
		IsSpent:     entry.IsSpent(),
		OutPoint: &corepb.OutPoint{
			Hash:  outPoint.Hash.GetBytes(),
			Index: outPoint.Index,
		},
		TxOut: &corepb.TxOut{
			Value:        entry.Value(),
			ScriptPubKey: entry.Script(),
		},
	}
}

func newMakeTxResp(
	code int32, msg string, tx *corepb.Transaction, rawMsgs [][]byte,
) *rpcpb.MakeTxResp {
	return &rpcpb.MakeTxResp{
		Code:    code,
		Message: msg,
		Tx:      tx,
		RawMsgs: rawMsgs,
	}
}

func (s *txServer) MakeUnsignedTx(
	ctx context.Context, req *rpcpb.MakeTxReq,
) (resp *rpcpb.MakeTxResp, err error) {

	logger.Infof("make unsigned tx: %+v", req)
	defer func() {
		if resp.Code != 0 {
			logger.Warnf("make unsigned tx %+v error: %s", req, resp.Message)
		} else {
			logger.Infof("make unsigned tx: %+v succeeded, response: %+v", resp)
		}
	}()
	wa := s.server.GetWalletAgent()
	if wa == nil || reflect.ValueOf(wa).IsNil() {
		return newMakeTxResp(-1, ErrAPINotSupported.Error(), nil, nil), nil
	}
	from, to := req.GetFrom(), req.GetTo()
	amounts, fee := req.GetAmounts(), req.GetFee()
	tx, utxos, err := rpcutil.MakeUnsignedTx(wa, from, to, amounts, fee)
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	pbTx, err := tx.ConvToPbTx()
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	rawMsgs, err := MakeTxRawMsgsForSign(tx, utxos...)
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	return newMakeTxResp(0, "", pbTx, rawMsgs), nil
}

func newMakeSplitAddrTxResp(code int32, msg string) *rpcpb.MakeSplitAddrTxResp {
	return &rpcpb.MakeSplitAddrTxResp{
		Code:    code,
		Message: msg,
	}
}

func (s *txServer) MakeUnsignedSplitAddrTx(
	ctx context.Context, req *rpcpb.MakeSplitAddrTxReq,
) (resp *rpcpb.MakeSplitAddrTxResp, err error) {

	logger.Infof("make unsigned split addr tx: %+v", req)
	defer func() {
		if resp.Code != 0 {
			logger.Warnf("make unsigned split addr tx %+v error: %s", req, resp.Message)
		} else {
			logger.Infof("make unsigned split addr tx: %+v succeeded, response: %+v", resp)
		}
	}()
	//
	wa := s.server.GetWalletAgent()
	if wa == nil || reflect.ValueOf(wa).IsNil() {
		return newMakeSplitAddrTxResp(-1, ErrAPINotSupported.Error()), nil
	}
	from, addrs := req.GetFrom(), req.GetAddrs()
	weights, fee := req.GetWeights(), req.GetFee()
	// make tx without sign
	tx, splitAddr, utxos, err := rpcutil.MakeUnsignedSplitAddrTx(wa, from, addrs,
		weights, fee)
	if err != nil {
		return newMakeSplitAddrTxResp(-1, err.Error()), nil
	}
	pbTx, err := tx.ConvToPbTx()
	if err != nil {
		return newMakeSplitAddrTxResp(-1, err.Error()), nil
	}
	// calc raw msgs
	rawMsgs, err := MakeTxRawMsgsForSign(tx, utxos...)
	if err != nil {
		return newMakeSplitAddrTxResp(-1, err.Error()), nil
	}
	resp = newMakeSplitAddrTxResp(0, "success")
	resp.SplitAddr, resp.Tx, resp.RawMsgs = splitAddr, pbTx, rawMsgs
	return resp, nil
}

func newMakeTokenIssueTxResp(code int32, msg string) *rpcpb.MakeTokenIssueTxResp {
	return &rpcpb.MakeTokenIssueTxResp{
		Code:    code,
		Message: msg,
	}
}

func (s *txServer) MakeUnsignedTokenIssueTx(
	ctx context.Context, req *rpcpb.MakeTokenIssueTxReq,
) (resp *rpcpb.MakeTokenIssueTxResp, err error) {

	logger.Infof("make unsigned token issue tx: %+v", req)
	defer func() {
		if resp.Code != 0 {
			logger.Warnf("make unsigned token issue tx %+v error: %s", req, resp.Message)
		} else {
			logger.Infof("make unsigned token issue tx: %+v succeeded, response: %+v", resp)
		}
	}()
	//
	wa := s.server.GetWalletAgent()
	if wa == nil || reflect.ValueOf(wa).IsNil() {
		return newMakeTokenIssueTxResp(-1, ErrAPINotSupported.Error()), nil
	}
	issuer, issuee, tag, fee := req.GetIssuer(), req.GetIssuee(), req.GetTag(), req.GetFee()
	// make tx without sign
	tx, issueOutIndex, utxos, err := rpcutil.MakeUnsignedTokenIssueTx(wa, issuer,
		issuee, tag, fee)
	if err != nil {
		return newMakeTokenIssueTxResp(-1, err.Error()), nil
	}
	pbTx, err := tx.ConvToPbTx()
	if err != nil {
		return newMakeTokenIssueTxResp(-1, err.Error()), nil
	}
	// calc raw msgs
	rawMsgs, err := MakeTxRawMsgsForSign(tx, utxos...)
	if err != nil {
		return newMakeTokenIssueTxResp(-1, err.Error()), nil
	}
	resp = newMakeTokenIssueTxResp(0, "success")
	resp.IssueOutIndex, resp.Tx, resp.RawMsgs = issueOutIndex, pbTx, rawMsgs
	return resp, nil
}

func (s *txServer) MakeUnsignedTokenTransferTx(
	ctx context.Context, req *rpcpb.MakeTokenTransferTxReq,
) (resp *rpcpb.MakeTxResp, err error) {

	logger.Infof("make unsigned token transfer tx: %+v", req)
	defer func() {
		if resp.Code != 0 {
			logger.Warnf("make unsigned token transfer tx %+v error: %s", req, resp.Message)
		} else {
			logger.Infof("make unsigned token transfer tx: %+v succeeded, response: %+v", resp)
		}
	}()
	wa := s.server.GetWalletAgent()
	if wa == nil || reflect.ValueOf(wa).IsNil() {
		return newMakeTxResp(-1, ErrAPINotSupported.Error(), nil, nil), nil
	}
	from, tid, fee := req.GetFrom(), req.GetTokenID(), req.GetFee()
	to, amounts := req.GetTo(), req.GetAmounts()
	tokenID := (*txlogic.TokenID)(txlogic.ConvPbOutPoint(tid))
	tx, utxos, err := rpcutil.MakeUnsignedTokenTransferTx(wa, from, to, amounts,
		tokenID, fee)
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	pbTx, err := tx.ConvToPbTx()
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	rawMsgs, err := MakeTxRawMsgsForSign(tx, utxos...)
	if err != nil {
		return newMakeTxResp(-1, err.Error(), nil, nil), nil
	}
	return newMakeTxResp(0, "", pbTx, rawMsgs), nil
}

// MakeTxRawMsgsForSign make tx raw msg for sign
func MakeTxRawMsgsForSign(tx *types.Transaction, utxos ...*rpcpb.Utxo) ([][]byte, error) {
	if len(tx.Vin) != len(utxos) {
		return nil, errors.New("invalid param")
	}
	msgs := make([][]byte, 0, len(tx.Vin))
	for i, u := range utxos {
		spk := u.GetTxOut().GetScriptPubKey()
		newTx := tx.Copy()
		for j, in := range newTx.Vin {
			if i != j {
				in.ScriptSig = nil
			} else {
				in.ScriptSig = spk
			}
		}
		data, err := newTx.Marshal()
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, data)
	}
	return msgs, nil
}

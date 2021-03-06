// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package txpool

import (
	"os"
	"testing"

	"github.com/BOXFoundation/boxd/boxd/eventbus"
	"github.com/BOXFoundation/boxd/core"
	"github.com/BOXFoundation/boxd/core/chain"
	"github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/p2p"
	"github.com/BOXFoundation/boxd/script"
	_ "github.com/BOXFoundation/boxd/storage/memdb" // init memdb
	"github.com/facebookgo/ensure"
	"github.com/jbenet/goprocess"
)

// test setup
var (
	proc        = goprocess.WithSignals(os.Interrupt)
	bus         = eventbus.New()
	txpool      = NewTransactionPool(proc, p2p.NewDummyPeer(), chain.NewTestBlockChain(), bus)
	chainHeight = uint32(0)

	txOutIdx = uint32(0)
	value    = uint64(1)

	privKey, pubKey, _ = crypto.NewKeyPair()
	addr, _            = types.NewAddressFromPubKey(pubKey)
	scriptAddr         = addr.Hash()
	scriptPubKey       = script.PayToPubKeyHashScript(scriptAddr)
	tx0, _             = chain.CreateCoinbaseTx(addr.Hash(), chainHeight)
)

// create a child tx spending parent tx's output
func createChildTx(parentTx *types.Transaction) *types.Transaction {
	outPoint := types.OutPoint{
		Hash:  *getTxHash(parentTx),
		Index: txOutIdx,
	}
	txIn := &types.TxIn{
		PrevOutPoint: outPoint,
		ScriptSig:    []byte{},
		Sequence:     0,
	}
	vIn := []*types.TxIn{
		txIn,
	}
	txOut := &corepb.TxOut{
		Value:        value,
		ScriptPubKey: *scriptPubKey,
	}
	vOut := []*corepb.TxOut{txOut}
	tx := &types.Transaction{
		Version:  1,
		Vin:      vIn,
		Vout:     vOut,
		Magic:    1,
		LockTime: 0,
	}

	// sign it
	for txInIdx, txIn := range tx.Vin {
		sigHash, err := script.CalcTxHashForSig(*scriptPubKey, tx, txInIdx)
		if err != nil {
			return nil
		}
		sig, err := crypto.Sign(privKey, sigHash)
		if err != nil {
			return nil
		}
		scriptSig := script.SignatureScript(sig, pubKey.Serialize())
		txIn.ScriptSig = *scriptSig

		// test to ensure
		if err = script.Validate(scriptSig, scriptPubKey, tx, txInIdx); err != nil {
			return nil
		}
	}
	return tx
}

func getTxHash(tx *types.Transaction) *crypto.HashType {
	txHash, _ := tx.TxHash()
	return txHash
}

func verifyProcessTx(t *testing.T, tx *types.Transaction, expectedErr error,
	isTransactionInPool, isOrphanInPool bool) {

	err := txpool.ProcessTx(tx, core.DefaultMode)
	ensure.DeepEqual(t, err, expectedErr)
	verifyTxInPool(t, tx, isTransactionInPool, isOrphanInPool)
}

func verifyTxInPool(t *testing.T, tx *types.Transaction, isTransactionInPool, isOrphanInPool bool) {
	txHash := getTxHash(tx)
	ensure.DeepEqual(t, isTransactionInPool, txpool.isTransactionInPool(txHash))
	ensure.DeepEqual(t, isOrphanInPool, txpool.isOrphanInPool(txHash))
}

func TestProcessTx(t *testing.T) {
	// Notations
	// txi <- txj: txj spends output of txi, i.e., is a child tx of txi
	// txi(m): txi is in main pool
	// txi(o): txi is in orphan pool
	// txi(): txi is in neither pool. This can happen, e.g., if a tx is a coinbase and rejected

	// tx0(): coinbase transaction cannot be accepted
	verifyProcessTx(t, tx0, core.ErrCoinbaseTx, false, false)

	// manually add tx0 into pool as utxo to bootstrap; otherwise no tx can be accepted
	txpool.addTx(tx0, chainHeight, 0)

	// tx0(m) <- tx1(m)
	// tx1 is admitted into main pool since it spends from a valid UTXO, i.e., coinbaseTx
	tx1 := createChildTx(tx0)
	verifyProcessTx(t, tx1, nil, true, false)

	// a duplicate tx1, already exists in tx pool
	verifyProcessTx(t, tx1, core.ErrDuplicateTxInPool, true, false)

	// tx0(m) <- tx1(m) <- tx2(m)
	tx2 := createChildTx(tx1)
	verifyProcessTx(t, tx2, nil, true, false)

	// tx2 is already in the main pool. Ignore.
	verifyProcessTx(t, tx2, core.ErrDuplicateTxInPool, true, false)

	// Withhold tx3 for now
	tx3 := createChildTx(tx2)

	// tx4 is orphaned since tx3 is missing
	// tx0(m) <- tx1(m) <- tx2(m)
	// tx4(o)
	tx4 := createChildTx(tx3)
	verifyProcessTx(t, tx4, core.ErrOrphanTransaction, false, true)

	// tx5 is orphaned since tx4 is missing
	// tx0(m) <- tx1(m) <- tx2(m)
	// tx4(o) <- tx5(o)
	tx5 := createChildTx(tx4)
	verifyProcessTx(t, tx5, core.ErrOrphanTransaction, false, true)

	// tx6 is orphaned since tx5 is missing
	// tx0(m) <- tx1(m) <- tx2(m)
	// tx4(o) <- tx5(o) <- tx6(o)
	tx6 := createChildTx(tx5)
	verifyProcessTx(t, tx6, core.ErrOrphanTransaction, false, true)

	ensure.DeepEqual(t, len(txpool.GetAllTxs()), 3)

	ensure.DeepEqual(t, len(txpool.GetOrphaTxs()), 3)

	// Add missing tx3
	verifyProcessTx(t, tx3, nil, true, false)

	// tx0(m) <- tx1(m) <- tx2(m) <- tx3(m) <- tx4(m) <- tx5(m) <- tx6(m)
	ensure.DeepEqual(t, len(txpool.GetAllTxs()), 7)

	// recursively remove tx4 and its children
	txpool.removeTx(tx4, true /* recursive */)
	// tx0(m) <- tx1(m) <- tx2(m) <- tx3(m)
	// tx5(o) <- tx6(o)
	ensure.DeepEqual(t, len(txpool.GetAllTxs()), 4)
	verifyTxInPool(t, tx4, false, false)
	verifyTxInPool(t, tx5, false, true)
	verifyTxInPool(t, tx6, false, true)

	// non-recursively remove tx1: its children remain in main pool
	txpool.removeTx(tx1, false /* recursive */)
	// tx0(m)
	// tx2(m) <- tx3(m)
	// tx5(o) <- tx6(o)
	ensure.DeepEqual(t, len(txpool.GetAllTxs()), 3)
	verifyTxInPool(t, tx1, false, false)
}

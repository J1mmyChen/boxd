// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chain

import (
	"os"

	"github.com/BOXFoundation/boxd/boxd/eventbus"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/p2p"
	"github.com/BOXFoundation/boxd/storage"
	"github.com/jbenet/goprocess"
	peer "github.com/libp2p/go-libp2p-peer"
)

// DummySyncManager is only used to test
type DummySyncManager struct{}

// NewDummySyncManager returns a new DummySyncManager
func NewDummySyncManager() *DummySyncManager {
	return &DummySyncManager{}
}

// StartSync starts sync
func (dm *DummySyncManager) StartSync() {}

// Run starts run
func (dm *DummySyncManager) Run() {}

// ActiveLightSync active light sync from remote peer.
func (dm *DummySyncManager) ActiveLightSync(pid peer.ID) error { return nil }

// NewTestBlockChain generate a chain for testing
func NewTestBlockChain() *BlockChain {
	dbCfg := &storage.Config{
		Name: "memdb",
		Path: "~/tmp",
	}

	proc := goprocess.WithSignals(os.Interrupt)
	db, _ := storage.NewDatabase(proc, dbCfg)
	blockChain, _ := NewBlockChain(proc, p2p.NewDummyPeer(), db, eventbus.Default())
	// set sync manager
	blockChain.Setup(new(DummyDpos), NewDummySyncManager())
	return blockChain
}

// DummyDpos dummy dpos
type DummyDpos struct{}

// Run dummy dpos
func (dpos *DummyDpos) Run() error { return nil }

// Stop dummy dpos
func (dpos *DummyDpos) Stop() {}

// RecoverMint revover mint
func (dpos *DummyDpos) RecoverMint() {}

// StopMint stop mint
func (dpos *DummyDpos) StopMint() {}

// Verify verify block.
func (dpos *DummyDpos) Verify(*types.Block) error { return nil }

// Process notify consensus to process new block.
func (dpos *DummyDpos) Process(*types.Block, interface{}) error { return nil }

// Seal notify consensus to change new tail.
func (dpos *DummyDpos) Seal(*types.Block) error { return nil }

// VerifyTx notify consensus to verify new tx.
func (dpos *DummyDpos) VerifyTx(*types.Transaction) error { return nil }

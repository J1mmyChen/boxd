// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wallet

import (
	"errors"
	"fmt"

	"github.com/BOXFoundation/boxd/boxd/eventbus"
	"github.com/BOXFoundation/boxd/core/chain"
	"github.com/BOXFoundation/boxd/log"
	"github.com/BOXFoundation/boxd/rpc/pb"
	"github.com/BOXFoundation/boxd/storage"
	"github.com/jbenet/goprocess"
)

var logger = log.NewLogger("wallet")

// Server is the struct type of an wallet service
type Server struct {
	proc  goprocess.Process
	bus   eventbus.Bus
	table storage.Table
	cfg   *Config
	//sync.Mutex
	//utxosQueue *list.List
}

// NewServer creates an Server instance using config and storage
func NewServer(parent goprocess.Process, config *Config, s storage.Storage,
	bus eventbus.Bus) (*Server, error) {
	proc := goprocess.WithParent(parent)
	table, err := s.Table(chain.BlockTableName)
	if err != nil {
		return nil, err
	}
	wServer := &Server{
		proc:  proc,
		bus:   bus,
		table: table,
		cfg:   config,
		//utxosQueue: list.New(),
	}
	return wServer, nil
}

// Run starts Server main loop
func (w *Server) Run() error {
	logger.Info("Wallet Server Start Running")
	//if err := w.initListener(); err != nil {
	//	return fmt.Errorf("fail to subscribe utxo change")
	//}
	//w.proc.Go(w.loop)
	return nil
}

func (w *Server) loop(p goprocess.Process) {
	for {
		// check whether to exit
		select {
		case <-p.Closing():
			logger.Infof("Quit Wallet Server")
			return
		default:
		}
		// process
		//elem := w.utxosQueue.Front()
		//if elem == nil {
		//	continue
		//}
		//value := w.utxosQueue.Remove(elem)
		//utxoSet := value.(*chain.UtxoSet)

		//allUtxos := utxoSet.All()
	}
}

// Proc returns then go process of the wallet server
func (w *Server) Proc() goprocess.Process {
	return w.proc
}

// Stop terminate the Server process
func (w *Server) Stop() {
	logger.Info("Wallet Server Stop Running")
	w.bus.Unsubscribe(eventbus.TopicUtxoUpdate, w.onUtxoChange)
	w.proc.Close()
}

func (w *Server) initListener() error {
	return w.bus.SubscribeAsync(eventbus.TopicUtxoUpdate, w.onUtxoChange, true)
}

func (w *Server) onUtxoChange(utxoSet *chain.UtxoSet) {
	//w.Lock()
	//defer w.Unlock()
	//w.utxosQueue.PushBack(utxoSet)
}

// Balance returns the total balance of an address
func (w *Server) Balance(addr string) (uint64, error) {
	if w.cfg == nil || !w.cfg.Enable {
		err := errors.New("not supported for non-wallet node")
		logger.Warn(err)
		return 0, err
	}
	balance, err := BalanceFor(addr, w.table)
	if err != nil {
		logger.Warnf("BalanceFor %s error %s", addr, err)
	}
	return balance, err
}

// Utxos returns all utxos of an address
func (w *Server) Utxos(addr string, amount uint64) ([]*rpcpb.Utxo, error) {
	if w.cfg == nil || !w.cfg.Enable {
		return nil, fmt.Errorf("fetch utxos not supported for non-wallet node")
	}
	utxos, err := FetchUtxosOf(addr, amount, w.table)
	if err != nil {
		logger.Warnf("Utxos for %s error %s", addr, err)
		return nil, err
	}
	return utxos, nil
}
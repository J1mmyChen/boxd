// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package walletserver

import (
<<<<<<< HEAD
	"container/list"
	"errors"
	"fmt"
	"sync"
=======
	"errors"
	"fmt"
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03

	"github.com/BOXFoundation/boxd/boxd/eventbus"
	"github.com/BOXFoundation/boxd/core/chain"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/log"
	"github.com/BOXFoundation/boxd/storage"
	"github.com/BOXFoundation/boxd/wallet/utxo"
	"github.com/jbenet/goprocess"
)

var logger = log.NewLogger("wallet-server")

// WalletServer is the struct type of an wallet service
type WalletServer struct {
	proc  goprocess.Process
	bus   eventbus.Bus
	table storage.Table
	cfg   *Config
<<<<<<< HEAD
	sync.Mutex
	utxosQueue *list.List
=======
	//sync.Mutex
	//utxosQueue *list.List
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
}

// NewWalletServer creates an WalletServer instance using config and storage
func NewWalletServer(parent goprocess.Process, config *Config, s storage.Storage,
	bus eventbus.Bus) (*WalletServer, error) {
	proc := goprocess.WithParent(parent)
	table, err := s.Table(chain.BlockTableName)
	if err != nil {
		return nil, err
	}
	wServer := &WalletServer{
<<<<<<< HEAD
		proc:       proc,
		bus:        bus,
		table:      table,
		cfg:        config,
		utxosQueue: list.New(),
=======
		proc:  proc,
		bus:   bus,
		table: table,
		cfg:   config,
		//utxosQueue: list.New(),
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
	}
	return wServer, nil
}

// Run starts WalletServer main loop
func (w *WalletServer) Run() error {
	logger.Info("Wallet Server Start Running")
<<<<<<< HEAD
	if err := w.initListener(); err != nil {
		return fmt.Errorf("fail to subscribe utxo change")
	}
=======
	//if err := w.initListener(); err != nil {
	//	return fmt.Errorf("fail to subscribe utxo change")
	//}
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
	//w.proc.Go(w.loop)
	return nil
}

func (w *WalletServer) loop(p goprocess.Process) {
	for {
		// check whether to exit
		select {
		case <-p.Closing():
			logger.Infof("Quit Wallet Server")
			return
		default:
		}
		// process
<<<<<<< HEAD
		elem := w.utxosQueue.Front()
		if elem == nil {
			continue
		}
		value := w.utxosQueue.Remove(elem)
		utxoSet := value.(*chain.UtxoSet)

		allUtxos := utxoSet.All()
		if err := utxo.ApplyUtxos(allUtxos, w.table); err != nil {
			logger.Warnf("wallet server fail to apply %d utxos error: %s",
				len(allUtxos), err)
			for op := range allUtxos {
				logger.Warnf("may unsucessfull applying utxo: %s", op)
			}
		}
=======
		//elem := w.utxosQueue.Front()
		//if elem == nil {
		//	continue
		//}
		//value := w.utxosQueue.Remove(elem)
		//utxoSet := value.(*chain.UtxoSet)

		//allUtxos := utxoSet.All()
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
	}
}

// Proc returns then go process of the wallet server
func (w *WalletServer) Proc() goprocess.Process {
	return w.proc
}

// Stop terminate the WalletServer process
func (w *WalletServer) Stop() {
	logger.Info("Wallet Server Stop Running")
	w.bus.Unsubscribe(eventbus.TopicUtxoUpdate, w.onUtxoChange)
	w.proc.Close()
}

func (w *WalletServer) initListener() error {
	return w.bus.SubscribeAsync(eventbus.TopicUtxoUpdate, w.onUtxoChange, true)
}

func (w *WalletServer) onUtxoChange(utxoSet *chain.UtxoSet) {
	//w.Lock()
	//defer w.Unlock()
	//w.utxosQueue.PushBack(utxoSet)
}

// Balance returns the total balance of an address
func (w *WalletServer) Balance(addr types.Address) (uint64, error) {
	if w.cfg == nil || !w.cfg.Enable {
		err := errors.New("not supported for non-wallet node")
		logger.Warn(err)
		return 0, err
<<<<<<< HEAD
	}
	balance, err := utxo.BalanceFor(addr, w.table)
	if err != nil {
		logger.Warnf("BalanceFor %s error %s", addr, err)
	}
=======
	}
	balance, err := utxo.BalanceFor(addr, w.table)
	if err != nil {
		logger.Warnf("BalanceFor %s error %s", addr, err)
	}
>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
	return balance, err
}

// Utxos returns all utxos of an address
<<<<<<< HEAD
func (w *WalletServer) Utxos(addr types.Address) (map[types.OutPoint]*types.UtxoWrap, error) {
=======
func (w *WalletServer) Utxos(addr types.Address) (types.UtxoMap, error) {

>>>>>>> 53cd91366c37fda6adeefee2ef94505631863d03
	if w.cfg == nil || !w.cfg.Enable {
		return nil, fmt.Errorf("fetch utxos not supported for non-wallet node")
	}
	return utxo.FetchUtxosOf(addr, w.table)
}

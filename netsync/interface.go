// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/yaochaoutokyo/bsvd/blockchain"
	"github.com/yaochaoutokyo/bsvd/chaincfg"
	"github.com/yaochaoutokyo/bsvd/chaincfg/chainhash"
	"github.com/yaochaoutokyo/bsvd/mempool"
	"github.com/yaochaoutokyo/bsvd/peer"
	"github.com/yaochaoutokyo/bsvd/wire"
	"github.com/yaochaoutokyo/bsvutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *bsvutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator

	MinSyncPeerNetworkSpeed uint64
}

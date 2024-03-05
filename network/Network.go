package network

import (
	"Go-Minichain/config"
	"Go-Minichain/consensus"
	"Go-Minichain/data"
)

type NetWork struct {
	blockchain data.BlockChain
	miner      consensus.MinerNode
}

func NewNetWork() *NetWork {
	blockchain := data.NewBlockChain()
	miner := consensus.NewMinerNode(data.NewTransactionPool(config.MiniChainConfig.GetMaxTransactionCount(), blockchain), blockchain)
	return &NetWork{
		blockchain: *blockchain,
		miner:      *miner,
	}
}
func (n *NetWork) Start() {
	n.miner.Run()
}

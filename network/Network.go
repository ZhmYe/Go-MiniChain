package network

import (
	"Go-Minichain/config"
	"Go-Minichain/data"
	"Go-Minichain/spv"
	"fmt"
)

type NetWork struct {
	accounts   []data.Account
	txPool     *TransactionPool
	blockchain *BlockChain
	miner      MinerNode
	spvPeer    []*SPVPeer
}

func NewNetWork() *NetWork {
	network := new(NetWork)
	fmt.Println("Accounts and SpvPeers config...")
	accounts := make([]data.Account, config.MiniChainConfig.GetAccountNumber())
	peers := make([]*SPVPeer, config.MiniChainConfig.GetAccountNumber())
	for i, _ := range accounts {
		accounts[i] = *data.NewAccount()
		peers[i] = NewSPVPeer(accounts[i], network) // 每个账户创建一个spv节点
		fmt.Println("	Network register new Account: ", accounts[i].GetWalletAddress())
	}
	network.accounts = accounts
	network.spvPeer = peers
	fmt.Println("Transaction Pool config...")
	pool := NewTransactionPool(config.MiniChainConfig.GetMaxTransactionCount(), network)
	fmt.Println("	Blockchain config...")
	blockchain := NewBlockChain(network)
	fmt.Println("Miner Peer config...")
	miner := NewMinerNode(network)
	fmt.Println("Network Config Finished...")
	fmt.Println("Network Start...")
	network.txPool = pool
	network.blockchain = blockchain
	network.miner = *miner
	// 广播创世块
	return network
}
func (n *NetWork) Start() {
	n.blockchain.SetUp()
	n.BroadCast(*n.GetNewestBlock())
	n.txPool.Start()
	n.miner.Run()
}
func (n *NetWork) CheckTransactionIsFull() bool {
	return n.txPool.IsFull()
}
func (n *NetWork) GetAllTransactions() []data.Transaction {
	return n.txPool.GetAll()
}
func (n *NetWork) GetTotalAmount() int {
	return n.blockchain.GetAllAmount()
}
func (n *NetWork) AddNewBlock(block data.Block) {
	n.blockchain.AddNewBlock(block)
}
func (n *NetWork) GetNewestBlock() *data.Block {
	return n.blockchain.GetNewestBlock()
}
func (n *NetWork) GetAccounts() []data.Account {
	return n.accounts
}
func (n *NetWork) GetAccount(i int) data.Account {
	return n.accounts[i]
}
func (n *NetWork) GetTrueUTXOs(address string) []*data.UTXO {
	return n.blockchain.GetTrueUTXOs(address)
}
func (n *NetWork) ProcessTransactionUTXO(inUTXO []*data.UTXO, outUTXO []*data.UTXO) {
	n.blockchain.ProcessTransactionUTXO(inUTXO, outUTXO)
}
func (n *NetWork) GetBlocks() []data.Block {
	return n.blockchain.GetBlocks()
}
func (n *NetWork) GetSPVPeers() []*SPVPeer {
	return n.spvPeer
}
func (n *NetWork) BroadCast(block data.Block) {
	n.miner.BroadCast(block)
}
func (n *NetWork) GetTransactionsInLatestBlock(address string) []data.Transaction {
	txs := make([]data.Transaction, 0)
	block := n.GetNewestBlock()
	blockBody := block.GetBlockBody()
	for _, transaction := range blockBody.GetTransctions() {
		have := false
		for _, utxo := range transaction.GetInUTXOs() {
			if utxo.GetWalletAddress() == address {
				txs = append(txs, transaction)
				have = true
				break
			}
		}
		if have {
			continue
		}
		for _, utxo := range transaction.GetOutUTXOs() {
			if utxo.GetWalletAddress() == address {
				txs = append(txs, transaction)
				break
			}
		}
	}
	return txs
}
func (n *NetWork) GetProof(hash string) spv.Proof {
	return n.miner.GetProof(hash)
}

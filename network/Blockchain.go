package network

import (
	"Go-Minichain/config"
	"Go-Minichain/data"
	"Go-Minichain/utils"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

/**
 * 区块链的类抽象，创建该对象时会自动生成创世纪块，加入区块链中
 */

type BlockChain struct {
	chain   []data.Block
	network *NetWork
	UTXOs   []*data.UTXO
	mutex   sync.Mutex
}

func NewBlockChain(network *NetWork) *BlockChain {
	chain := new(BlockChain)
	chain.chain = make([]data.Block, 0)
	chain.UTXOs = make([]*data.UTXO, 0)
	chain.network = network
	return chain
}
func (c *BlockChain) SetUp() {
	transactions := c.GenesisTransactions()
	header := data.NewBlockHeader("", "", rand.Int63())
	body := c.network.miner.GetBlockBody(transactions)
	genesisBlock := data.NewBlock(*header, body)
	fmt.Println("Create the genesis Block! ")
	fmt.Println("And the hash of genesis Block is : " + fmt.Sprint(utils.Sha256Digest([]byte(genesisBlock.ToString()))) +
		", you will see the hash value in next Block's preBlockHash field.")
	fmt.Println()
	c.AddNewBlock(*genesisBlock)
}
func (c *BlockChain) AddNewBlock(block data.Block) {
	c.chain = append(c.chain, block)
}
func (c *BlockChain) GetNewestBlock() *data.Block {
	return &c.chain[len(c.chain)-1]
}
func (c *BlockChain) GenesisTransactions() []data.Transaction {
	outUTXOs := make([]*data.UTXO, len(c.network.GetAccounts()))
	for i := 0; i < len(outUTXOs); i++ {
		account := c.network.GetAccount(i)
		outUTXOs[i] = data.NewUTXO(account.GetWalletAddress(), config.MiniChainConfig.GetInitAmount(), account.GetPublicKey())
	}
	c.ProcessTransactionUTXO([]*data.UTXO{}, outUTXOs)
	daydreamPrivateKey, daydreamPublicKey := utils.Secp256k1Generate()
	sign := utils.Signature([]byte("Wecome to Blockchain Lab!!!"), daydreamPrivateKey)
	return []data.Transaction{*data.NewTransaction(make([]*data.UTXO, 0), outUTXOs, sign, daydreamPublicKey)}
}
func (c *BlockChain) AddUTXO(u *data.UTXO) {
	c.UTXOs = append(c.UTXOs, u)
}

/**
 * 遍历整个区块链获得某钱包地址相关的utxo，获得真正的utxo，即未被使用的utxo
 * @param walletAddress 钱包地址
 * @return
 */

func (c *BlockChain) GetTrueUTXOs(walletAddress string) []*data.UTXO {
	trueUTXOs := make([]*data.UTXO, 0)
	// 这里直接遍历utxos
	for _, utxo := range c.UTXOs {
		if !utxo.IsUsed() && utxo.GetWalletAddress() == walletAddress {
			trueUTXOs = append(trueUTXOs, utxo)
		}
	}
	return trueUTXOs
}

//	func (c *BlockChain) GetAccount() []data.Account {
//		return c.accounts
//	}
func (c *BlockChain) GetAllAmount() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	sumAccount := 0
	accounts := c.network.GetAccounts()
	for _, account := range accounts {
		utxos := c.GetTrueUTXOs(account.GetWalletAddress())
		for _, utxo := range utxos {
			if utxo.IsUsed() {
				panic("error")
			}
		}
		sumAccount += account.GetAmount(utxos)
	}
	// 也可以写成
	//for _, utxo := range c.UTXOs {
	//	if !utxo.IsUsed() {
	//		sumAccount += utxo.GetAmount()
	//	}
	//}
	if sumAccount != config.MiniChainConfig.GetAccountNumber()*config.MiniChainConfig.GetInitAmount() {
		panic("error Balance:" + strconv.Itoa(sumAccount))
	}
	return sumAccount
}
func (c *BlockChain) ProcessTransactionUTXO(inUTXOs []*data.UTXO, outUTXO []*data.UTXO) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, utxo := range inUTXOs {
		utxo.SetUsed()
	}
	for _, utxo := range outUTXO {
		c.AddUTXO(utxo)
	}
}

func (c *BlockChain) GetBlocks() []data.Block {
	return c.chain
}

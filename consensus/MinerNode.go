package consensus

import (
	"Go-Minichain/config"
	"Go-Minichain/data"
	"Go-Minichain/utils"
	"fmt"
	"math/rand"
	"strings"
)

/**
 * 矿工线程
 *
 * 该线程的主要工作就是不断的进行交易打包、Merkle树根哈希值计算、构造区块，
 * 然后尝试使用不同的随机字段（nonce）进行区块的哈希值计算以生成新的区块添加到区块中
 *
 * 这里需要你实现的功能函数为：getBlockBody、mine和getBlock，具体的需求见上述方法前的注释，
 * 除此之外，该类中的其他方法、变量，以及其他类中的方法和变量，均无需修改，否则可能影响系统的正确运行
 *
 * 如有疑问，及时交流
 *
 */

type MinerNode struct {
	transactionPool *data.TransactionPool
	blockchain      *data.BlockChain
}

func NewMinerNode(pool *data.TransactionPool, chain *data.BlockChain) *MinerNode {
	return &MinerNode{transactionPool: pool, blockchain: chain}
}
func (m *MinerNode) Run() {
	m.transactionPool.Start()
	for {
		if m.transactionPool.IsFull() {
			transactions := m.transactionPool.GetAll()
			blockBody := m.GetBlockBody(transactions)
			m.Mine(blockBody)
			fmt.Println("The Sum of all amount: ", m.blockchain.GetAllAmount())
		}
	}
}
func (m *MinerNode) GetBlockBody(transactions []data.Transaction) data.BlockBody {
	if transactions == nil || len(transactions) > config.MiniChainConfig.GetMaxTransactionCount() {
		panic("transactions can not be nil or be more than config.MaxTransactionCount")
	}
	// todo 这里计算merkle Root
	hashs := make([]string, 0)
	for _, transaction := range transactions {
		txHash := utils.GetSha256Digest(transaction.ToString())
		hashs = append(hashs, txHash)
	}
	for {
		if len(hashs) == 1 {
			break
		}
		nbHashs := len(hashs)
		for i := 0; i < nbHashs; i += 2 {
			leftNode := hashs[i]
			rightNode := hashs[i+1]
			hashs = append(hashs, utils.GetSha256Digest(leftNode+rightNode))
		}
		if nbHashs%2 == 1 {
			node := hashs[nbHashs-1]
			hashs = append(hashs, utils.GetSha256Digest(node+node))
		}
		hashs = hashs[nbHashs:]
	}
	return *data.NewBlockBody(hashs[0], transactions)
}

/**
 * 该方法供mine方法调用，其功能为根据传入的区块体参数，构造一个区块对象返回，
 * 也就是说，你需要构造一个区块头对象，然后用一个区块对象组合区块头和区块体
 *
 * 建议查看BlockHeader类中的字段和注释，有助于你实现该方法
 *
 * @param blockBody 区块体
 *
 * @return 相应的区块对象
 */

func (m *MinerNode) Mine(blockBody data.BlockBody) {
	block := m.GetBlock(blockBody)
	for {
		blockHash := utils.GetSha256Digest(block.ToString())
		if strings.HasPrefix(blockHash, utils.HashPrefixTarget()) {
			header := block.GetBlockHeader()
			fmt.Println("Mined a new Block! Previous Block Hash is: " + header.GetPreBlockHash())
			fmt.Println("And the hash of this Block is : " + utils.GetSha256Digest(block.ToString()) +
				", you will see the hash value in next Block's preBlockHash field.")
			fmt.Println()
			m.blockchain.AddNewBlock(*block)
			break
		} else {
			// todo
			nonce := rand.Int63()
			block.SetNonce(nonce)
		}
	}
}
func (m *MinerNode) GetBlock(blockBody data.BlockBody) *data.Block {
	// todo
	preBlock := m.blockchain.GetNewestBlock()
	preBlockHash := utils.GetSha256Digest(preBlock.ToString())
	nonce := rand.Int63()
	blockHeader := data.NewBlockHeader(preBlockHash, blockBody.GetMerkleRootHash(), nonce)
	block := data.NewBlock(*blockHeader, blockBody)
	return block
}

func (m *MinerNode) Check(transactions []data.Transaction) bool {
	for _, transaction := range transactions {
		data := data.UTXO2Bytes(transaction.GetInUTXOs(), transaction.GetOutUTXOs())
		sign := transaction.GetSendSign()
		publicKey := transaction.GetSendPublicKey()
		if !utils.Verify(data, sign, &publicKey) {
			return false
		}
	}
	return true
}

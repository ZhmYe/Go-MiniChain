package data

import (
	"Go-Minichain/utils"
	"fmt"
	"math/rand"
)

/**
 * 区块链的类抽象，创建该对象时会自动生成创世纪块，加入区块链中
 */

type BlockChain struct {
	chain []Block
}

func NewBlockChain() *BlockChain {
	chain := new(BlockChain)
	chain.chain = make([]Block, 0)
	header := NewBlockHeader("", "", rand.Int63())
	body := NewBlockBody("", []Transaction{})
	genesisBlock := NewBlock(*header, *body)
	fmt.Println("Create the genesis Block! ")
	fmt.Println("And the hash of genesis Block is : " + utils.GetSha256Digest(genesisBlock.ToString()) +
		", you will see the hash value in next Block's preBlockHash field.")
	fmt.Println()
	chain.AddNewBlock(*genesisBlock)
	return chain
}
func (c *BlockChain) AddNewBlock(block Block) {
	c.chain = append(c.chain, block)
}
func (c *BlockChain) GetNewestBlock() *Block {
	return &c.chain[len(c.chain)-1]
}

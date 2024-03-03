package data

import (
	"fmt"
	"strings"
)

/**
 * 对区块体的抽象，主要有两个字段：
 *    transactions: 从交易池中取得的一批次交易
 *
 *    merkleRootHash: 使用上述交易，计算得到的Merkle树根哈希值
 */

type BlockBody struct {
	transactions   []Transaction
	merkleRootHash string
}

func NewBlockBody(merkleRootHash string, transactions []Transaction) *BlockBody {
	return &BlockBody{
		transactions:   transactions,
		merkleRootHash: merkleRootHash,
	}
}
func (b *BlockBody) GetTransctions() []Transaction {
	return b.transactions
}
func (b *BlockBody) GetMerkleRootHash() string {
	return b.merkleRootHash
}

func (b *BlockBody) toString() string {
	return "BlockBody{" +
		"merkleRootHash=" + b.merkleRootHash +
		", transactions=" + strings.Join([]string{"", fmt.Sprint(b.transactions)}, " ") +
		"}"
}

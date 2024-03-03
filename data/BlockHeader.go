package data

import (
	"Go-Minichain/config"
	"strconv"
	"time"
)

/**
 * 对区块头的抽象（参考比特币中的区块头结构），主要有以下字段：
 *    version: 版本号，默认为1，无需提供该参数
 *
 *    preBlockHash: 前一个区块的哈希值，创建新的区块头对象时需要提供该参数
 *
 *    merkleRootHash: 该区块头对应区块体中的交易的Merkle根哈希值，创建新的区块头对象时需要提供该参数
 *
 *    timestamp: 时间戳，创建区块头对象时会自动填充，无需提供该参数
 *
 *    difficulty: 挖矿难度，默认为系统配置中的难度值，无需提供该参数
 *
 *    nonce: 随机字段，创建新的区块头对象时需要提供该参数
 *
 */

type BlockHeader struct {
	version        int
	preBlockHash   string
	merkleRootHash string
	timestamp      int
	difficulty     int
	nonce          int64
}

func NewBlockHeader(preBlockHash string, merkleRootHash string, nonce int64) *BlockHeader {
	header := new(BlockHeader)
	header.version = 1
	header.timestamp = time.Now().Nanosecond()
	header.preBlockHash = preBlockHash
	header.difficulty = config.MiniChainConfig.GetDifficulty()
	header.nonce = nonce
	header.merkleRootHash = merkleRootHash
	return header
}
func (h *BlockHeader) GetVersion() int {
	return h.version
}
func (h *BlockHeader) GetPreBlockHash() string {
	return h.preBlockHash
}
func (h *BlockHeader) GetMerkleRootHash() string {
	return h.merkleRootHash
}
func (h *BlockHeader) GetTimestamp() int {
	return h.timestamp
}
func (h *BlockHeader) GetDifficulty() int {
	return h.difficulty
}
func (h *BlockHeader) GetNonce() int64 {
	return h.nonce
}

func (h *BlockHeader) SetNonce(nonce int64) {
	h.nonce = nonce
}

func (h *BlockHeader) toString() string {
	return "BlockHeader{" +
		"version=" + strconv.Itoa(h.version) +
		", preBlockHash=" + h.preBlockHash +
		", merkleRootHash=" + h.merkleRootHash +
		", timeStamp=" + strconv.Itoa(h.timestamp) +
		", difficulty=" + strconv.Itoa(h.difficulty) +
		", nonce=" + strconv.FormatInt(h.nonce, 10) +
		"}"

}

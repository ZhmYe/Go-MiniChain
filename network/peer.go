package network

import (
	"Go-Minichain/data"
	"Go-Minichain/spv"
	"Go-Minichain/utils"
)

type SPVPeer struct {
	headers []data.BlockHeader // SPV节点只存储区块头
	account data.Account       // spv节点的账户信息
	network *NetWork           // 节点连入网络
}

func NewSPVPeer(account data.Account, network *NetWork) *SPVPeer {
	return &SPVPeer{
		headers: []data.BlockHeader{},
		account: account,
		network: network,
	}
}

/**
 * 添加一个区块头
 *
 */

func (p *SPVPeer) Accept(header data.BlockHeader) {
	p.headers = append(p.headers, header)
	if len(p.network.GetBlocks()) == 1 {
		// 创世块不需要验证
		return
	}
	if !p.VerifyHeader() {
		panic("SPV Verify Not Pass!!!")
	}
}
func (p *SPVPeer) Verify(transaction data.Transaction) bool {
	// 获取交易哈希
	txHash := utils.GetSha256Digest(transaction.ToString())
	// 通过网络得到proof
	proof := p.network.GetProof(txHash)
	hash := proof.GetTxHash()
	// 使用验证路径得到Merkle Root Hash
	for _, node := range proof.GetPath() {
		switch node.GetOrientation() {
		case spv.LEFT:
			hash = utils.GetSha256Digest(node.GetTxHash() + hash)
		case spv.RIGHT:
			hash = utils.GetSha256Digest(hash + node.GetTxHash())
		default:
			return false
		}
	}
	// 获得本地区块头的根哈希
	height := proof.GetHeight()
	localMerkleRootHash := p.headers[height].GetMerkleRootHash()

	//获得proof中的根哈希
	merleRootHash := proof.GetMerkleRootHash()
	// 这里可以输出看看结果
	//fmt.Println("\n-------------> Verify Hash: ", txHash)
	//fmt.Println("CalTxHash:", hash)
	//fmt.Println("LocalMerkleRootHash: ", localMerkleRootHash)
	//fmt.Println("RemoteMerkleRootHash: ", merleRootHash)

	return hash == localMerkleRootHash && hash == merleRootHash
}
func (p *SPVPeer) VerifyHeader() bool {
	transactions := p.network.GetTransactionsInLatestBlock(p.account.GetWalletAddress())
	if len(transactions) == 0 {
		return true
	}
	//fmt.Println("Account[", p.account.GetWalletAddress(), "] began to verify the transaction...")
	for _, transaction := range transactions {
		if !p.Verify(transaction) {
			return false
		}
	}
	return true
}

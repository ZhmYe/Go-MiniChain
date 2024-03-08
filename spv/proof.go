package spv

type Proof struct {
	txHash         string // 待验证交易的交易哈希
	merkleRootHash string // Merkle根哈希
	height         int    // 待验证交易所在区块高度
	path           []Node // 验证路径，内部是哈希值和偏向
}

func NewProof(hash string, merkleRootHash string, height int, path []Node) *Proof {
	return &Proof{
		txHash:         hash,
		merkleRootHash: merkleRootHash,
		height:         height,
		path:           path,
	}
}
func (p *Proof) GetTxHash() string {
	return p.txHash
}
func (p *Proof) GetMerkleRootHash() string {
	return p.merkleRootHash
}
func (p *Proof) GetHeight() int {
	return p.height
}
func (p *Proof) GetPath() []Node {
	return p.path
}

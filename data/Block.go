package data

/**
 * 区块的类抽象，组合了区块头和区块体
 *
 */

type Block struct {
	header BlockHeader // 区块头
	body   BlockBody   // 区块体
}

func NewBlock(header BlockHeader, body BlockBody) *Block {
	return &Block{header: header, body: body}
}
func (b *Block) GetBlockHeader() BlockHeader {
	return b.header
}
func (b *Block) GetBlockBody() BlockBody {
	return b.body
}
func (b *Block) ToString() string {
	return "Block{" +
		"blockHeader=" + b.header.toString() +
		", blockBody=" + b.body.toString() +
		"}"
}
func (b *Block) SetNonce(nonce int64) {
	b.header.SetNonce(nonce)
}

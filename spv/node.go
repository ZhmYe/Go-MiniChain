package spv

// Orientation 哈希在Merkle树上中作为节点的偏向，左节点/右节点
type Orientation int

const (
	LEFT Orientation = iota
	RIGHT
)

type Node struct {
	txHash      string
	orientation Orientation
}

func NewNode(hash string, o Orientation) *Node {
	return &Node{
		txHash:      hash,
		orientation: o,
	}
}
func (n *Node) GetTxHash() string {
	return n.txHash
}
func (n *Node) GetOrientation() Orientation {
	return n.orientation
}

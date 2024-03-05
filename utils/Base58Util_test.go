package utils

import (
	"fmt"
	"testing"
)

func TestBase58Util(t *testing.T) {
	b := NewBase58Util()
	data := []byte{0, 0, 58, 0, 0, 59}
	dataStr := b.Encode(data)
	fmt.Println("编码后：", dataStr)
	ndata := b.Decode(dataStr)
	fmt.Println(ndata)
}

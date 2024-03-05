package data

import (
	"Go-Minichain/utils"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"strconv"
)

type UTXO struct {
	walletAddress string
	amount        int
	publicKeyHash []byte
	used          bool // 是否已经被使用
}

func NewUTXO(address string, amount int, publicKey ecdsa.PublicKey) *UTXO {
	publicKeyBytes := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	return &UTXO{
		walletAddress: address,
		amount:        amount,
		publicKeyHash: utils.Ripemd160Digest(utils.Sha256Digest(publicKeyBytes)),
		used:          false,
	}
}
func (utxo *UTXO) SetUsed() {
	if utxo.used {
		panic("can't use an UTXO twice!!!")
	}
	utxo.used = true
}
func (utxo *UTXO) IsUsed() bool {
	return utxo.used
}

/**
 * 模拟utxo的解锁脚本，只有使用对应的私钥签名和公钥，正确解锁才能使用该utxo作为交易输入
 * @param sign 账户私钥签名，这里我们这么约定:签名数据为公钥二进制数据
 * @param publicKey 公钥
 * @return
 */

func (utxo *UTXO) UnlockScript(sign []byte, publicKey ecdsa.PublicKey) bool {
	stack := make([][]byte, 0)
	// <sig> 签名入栈
	// 栈内: <Sig>
	stack = append(stack, sign)
	// <PubK> 公钥入栈
	// 栈内: <Sig> <PubK>
	publicKeyBytes := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	stack = append(stack, publicKeyBytes)
	// DUP 复制一份栈顶数据, peek()为java栈容器获取栈顶元素的函数
	// 栈内: <Sig> <PubK> <PubK>
	stack = append(stack, stack[len(stack)-1])
	// HASH160 弹出栈顶元素，进行哈希摘要: RIPEMD160(SHA256(PubK)，并将其入栈
	// 栈内: <Sig> <PubK> <PubHash>
	data := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	stack = append(stack, utils.Ripemd160Digest(utils.Sha256Digest(data)))
	// <PubHash> utxo先前保存的公钥哈希入栈
	// 栈内: <Sig> <PubK> <PubHash> <PubHash>
	stack = append(stack, utxo.publicKeyHash)
	// EQUALVERIFY 比较栈顶的两个公钥哈希是否相同，不相同则解锁失败
	// 栈内: <Sig> <PubK>
	publicKeyHash1 := stack[len(stack)-1]
	publicKeyHash2 := stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	if !bytes.Equal(publicKeyHash1, publicKeyHash2) {
		return false
	}
	// CHECKSIG 检查签名是否正确，正确则入栈 TRUE;
	// 栈内:
	publicKeyEncoded := stack[len(stack)-1]
	sign1 := stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	// 比特币网络中因为其脚本支持操作少的特性，需要入栈再检查，这里验证正确我们就直接返回了
	// 栈内: TRUE (验证正确情况下）
	return utils.Verify(publicKeyEncoded, sign1, &publicKey)
}
func (utxo *UTXO) GetWalletAddress() string {
	return utxo.walletAddress
}
func (utxo *UTXO) GetAmount() int {
	return utxo.amount
}
func (utxo *UTXO) GetPublicKeyHash() []byte {
	return utxo.publicKeyHash
}
func (utxo *UTXO) ToString() string {
	return "UTXO{" +
		"walletAddress=" + utxo.walletAddress +
		", amount=" + strconv.Itoa(utxo.amount) +
		", publicKeyHash=" + utils.Byte2HexString(utxo.publicKeyHash) +
		"}"
}

/**
 * utxo数组（包含输入和输出）转化为byte数据供交易签名
 * @param inUtxos
 * @param outUtxos
 * @return
 */

func UTXO2Bytes(inUTXO []*UTXO, outUTXO []*UTXO) []byte {
	return []byte(fmt.Sprint(inUTXO) + fmt.Sprint(outUTXO))
}

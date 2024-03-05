package data

import (
	"Go-Minichain/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"strconv"
	"time"
)

/**
 * 对交易的抽象
 */

type Transaction struct {
	//data      string
	timestamp     int
	inUTXO        []*UTXO
	outUTXO       []*UTXO
	sendSign      []byte          // 交易发送方的私钥签名
	sendPublicKey ecdsa.PublicKey //交易发送方的公钥
}

func NewTransaction(inUTXO []*UTXO, outUTXO []*UTXO, sendSign []byte, sendPublicKey ecdsa.PublicKey) *Transaction {
	return &Transaction{
		inUTXO:        inUTXO,
		outUTXO:       outUTXO,
		sendSign:      sendSign,
		sendPublicKey: sendPublicKey,
		timestamp:     time.Now().Nanosecond(),
	}
}
func (t *Transaction) GetInUTXOs() []*UTXO {
	return t.inUTXO
}
func (t *Transaction) GetOutUTXOs() []*UTXO {
	return t.outUTXO
}
func (t *Transaction) GetSendSign() []byte {
	return t.sendSign
}
func (t *Transaction) GetSendPublicKey() ecdsa.PublicKey {
	return t.sendPublicKey
}

func (t *Transaction) GetTimeStamp() int {
	return t.timestamp
}

func (t *Transaction) ToString() string {
	publicKey := t.sendPublicKey
	publicKeyBytes := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	return "\nTransaction{" +
		"\ninUTXOs=" + fmt.Sprint(t.inUTXO) +
		", \noutUTXOs=" + fmt.Sprint(t.outUTXO) +
		", \nsendSign=" + utils.Byte2HexString(t.sendSign) +
		", \nsendPublicKey=" + utils.Byte2HexString(publicKeyBytes) +
		", timestamp=" + strconv.Itoa(t.timestamp) +
		"}"
}

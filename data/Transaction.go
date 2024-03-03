package data

import (
	"strconv"
	"time"
)

/**
 * 对交易的抽象
 */

type Transaction struct {
	data      string
	timestamp int
}

func NewTransaction(data string) Transaction {
	return Transaction{
		data:      data,
		timestamp: time.Now().Nanosecond(),
	}
}
func (t *Transaction) GetData() string {
	return t.data
}
func (t *Transaction) GetTimeStamp() int {
	return t.timestamp
}

func (t *Transaction) ToString() string {
	return "Transaction{" +
		"data=" + t.data +
		", timestamp=" + strconv.Itoa(t.timestamp) +
		"}"
}

package data

import (
	"Go-Minichain/utils"
	"math/rand"
)

/**
 * 交易池
 */

type TransactionPool struct {
	transactions []Transaction
	capacity     int
}

func NewTransactionPool(c int) *TransactionPool {
	p := new(TransactionPool)
	p.capacity = c
	p.transactions = make([]Transaction, 0)
	return p
}
func (p *TransactionPool) Put(transaction Transaction) {
	p.transactions = append(p.transactions, transaction)
}
func (p *TransactionPool) GetAll() []Transaction {
	transactions := p.transactions
	p.clear()
	return transactions
}
func (p *TransactionPool) clear() {
	p.transactions = make([]Transaction, 0)
}
func (p *TransactionPool) IsFull() bool {
	return len(p.transactions) >= p.capacity
}
func (p *TransactionPool) IsEmpty() bool {
	return len(p.transactions) == 0
}
func (p *TransactionPool) GetCapacity() int {
	return p.capacity
}
func (p *TransactionPool) Start() {
	go func() {
		for {
			if !p.IsFull() {
				transaction := NewTransaction(utils.RandomString(rand.Intn(10)))
				p.Put(transaction)
			}
		}
	}()
}

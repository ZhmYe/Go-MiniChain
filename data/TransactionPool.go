package data

import (
	"Go-Minichain/utils"
	"crypto/elliptic"
	"math/rand"
)

/**
 * 交易池
 */

type TransactionPool struct {
	transactions []Transaction
	capacity     int
	blockchain   *BlockChain
}

func NewTransactionPool(c int, chain *BlockChain) *TransactionPool {
	p := new(TransactionPool)
	p.capacity = c
	p.transactions = make([]Transaction, 0)
	p.blockchain = chain
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
func (p *TransactionPool) GetNewTransaction() *Transaction {
	accounts := p.blockchain.GetAccount()
	var transaction *Transaction
	for {
		// 随机获取两个账户A和B
		aAccount := accounts[rand.Intn(len(accounts))]
		bAccount := accounts[rand.Intn(len(accounts))]
		// BTC不允许自己给自己转账
		if aAccount == bAccount {
			continue
		}
		// 获得钱包地址
		aWalletAddress := aAccount.GetWalletAddress()
		bWalletAddress := bAccount.GetWalletAddress()
		// 获取A可用的Utxo并计算余额
		aTrueUtxos := p.blockchain.GetTrueUTXOs(aWalletAddress)
		aAmount := aAccount.GetAmount(aTrueUtxos)
		// 如果A账户的余额为0，则无法构建交易，重新随机生成
		if aAmount == 0 {
			continue
		}
		// 随机生成交易数额 [1, aAmount] 之间
		txAmount := rand.Intn(aAmount) + 1
		inUTXOs := make([]*UTXO, 0)
		outUTXOs := make([]*UTXO, 0)
		// A账户需先解锁才能使用自己的utxo，解锁需要私钥签名和公钥去执行解锁脚本，这里先生成需要解锁的签名
		// 签名的数据我们约定为公钥的二进制数据
		publickKey := aAccount.GetPublicKey()
		publicKeyBytes := elliptic.Marshal(publickKey, publickKey.X, publickKey.Y)
		aUnLockSign := utils.Signature(publicKeyBytes, aAccount.GetPrivateKey())
		//fmt.Println(utils.Verify(publicKeyBytes, aUnLockSign, &publickKey))
		// 选择输入总额>=交易数额的 utxo

		inAmount := 0
		for _, utxo := range aTrueUtxos {
			if utxo.UnlockScript(aUnLockSign, aAccount.GetPublicKey()) {
				inAmount += utxo.GetAmount()
				inUTXOs = append(inUTXOs, utxo)
				if inAmount >= txAmount {
					break
				}
			}
		}
		// 可解锁的utxo总额仍不足以支付交易数额，则重新随机
		if inAmount < txAmount {
			continue
		}
		// 构建输出OutUtxos，A账户向B账户支付txAmount，同时输入对方的公钥以供生成公钥哈希
		outUTXOs = append(outUTXOs, NewUTXO(bWalletAddress, txAmount, bAccount.GetPublicKey()))
		// 如果有余额，则“找零”，即给自己的utxo
		if inAmount > txAmount {
			outUTXOs = append(outUTXOs, NewUTXO(aWalletAddress, inAmount-txAmount, aAccount.GetPublicKey()))
		}
		// A账户需对整个交易进行私钥签名，确保交易不会被篡改，因为交易会传输到网络中，而上述步骤可在本地离线环境中构造
		// 获取要签名的数据: inUtxos & outUtxos
		data := UTXO2Bytes(inUTXOs, outUTXOs)
		// A账户使用私钥签名
		sign := utils.Signature(data, aAccount.GetPrivateKey())
		//timeStamp := time.Now().Nanosecond()
		transaction = NewTransaction(inUTXOs, outUTXOs, sign, aAccount.GetPublicKey())
		// 这里还需要将outUXTO添加,并将inUTXO置为已使用
		p.blockchain.ProcessTransactionUTXO(inUTXOs, outUTXOs)
		break
	}
	return transaction
}
func (p *TransactionPool) Start() {
	go func() {
		for {
			if !p.IsFull() {
				transaction := p.GetNewTransaction()
				p.Put(*transaction)
			}
		}
	}()
}

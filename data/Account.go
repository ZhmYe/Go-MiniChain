package data

import (
	"Go-Minichain/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
)

type Account struct {
	publicKey  ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

func NewAccount() *Account {
	privateKey, publicKey := utils.Secp256k1Generate()
	return &Account{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

/**
 * 根据账户的公钥计算钱包地址
 * @return
 */

func (a *Account) GetWalletAddress() string {
	// 公钥哈希：RIPEMD160(SHA256(PubK)
	publicKey := a.GetPublicKey()
	publicKeyBytes := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	publicKeyHash := utils.Ripemd160Digest(utils.Sha256Digest(publicKeyBytes))
	// 0x0 + 公钥哈希
	data := make([]byte, 1+len(publicKeyHash))
	data = append(data, 0)
	data = append(data, publicKeyHash...)
	doubleHash := utils.Sha256Digest(utils.Sha256Digest(data))
	// 0x0 + 公钥哈希 + 校验（两次哈希后前4字节）
	wallEncoded := make([]byte, 1+len(publicKeyHash)+4)
	wallEncoded = append(wallEncoded, 0)
	wallEncoded = append(wallEncoded, publicKeyHash...)
	wallEncoded = append(wallEncoded, doubleHash[:4]...)
	// 对二进制地址进行BASE58编码，得到钱包地址（字符串形式）
	b := utils.NewBase58Util()
	walletAddress := b.Encode(wallEncoded)
	return walletAddress
}

/**
 * 根据未使用的utxo计算账户的余额
 * @param trueUtxos 未使用的utxo
 * @return
 */

func (a *Account) GetAmount(trueUtxo []*UTXO) int {
	amount := 0
	for _, utxo := range trueUtxo {
		amount += utxo.GetAmount()
	}
	return amount
}
func (a *Account) ToString() string {
	publicKeyBytes, err := asn1.Marshal(a.publicKey)
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := asn1.Marshal(a.privateKey)
	if err != nil {
		panic(err)
	}
	return "Account{" +
		"publicKey=" + utils.Byte2HexString(publicKeyBytes) +
		"privateKey=" + utils.Byte2HexString(privateKeyBytes) +
		"}"
}
func (a *Account) GetPublicKey() ecdsa.PublicKey {
	return a.publicKey
}
func (a *Account) GetPrivateKey() *ecdsa.PrivateKey {
	return a.privateKey
}

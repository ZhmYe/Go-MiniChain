package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/dustinxie/ecc"
)

/**
 * 比特数据转为相应的十六进制字符串
 * @param data
 * @return
 */

func Byte2HexString(data []byte) string {
	return fmt.Sprint(data)
}

/**
 * secp256k1密钥生成
 * @return
 */

func Secp256k1Generate() (*ecdsa.PrivateKey, ecdsa.PublicKey) {
	// generate secp256k1 private key
	p256k1 := ecc.P256k1()
	privateKey, err := ecdsa.GenerateKey(p256k1, rand.Reader)
	if err != nil {
		panic("KeyPair Generate Error...")
	}
	return privateKey, privateKey.PublicKey
}

/**
 * 私钥签名
 * @param data 签名数据
 * @param privateKey 签名私钥
 * @return 签名后的比特数据
 */

func Signature(data []byte, privateKey *ecdsa.PrivateKey) []byte {
	// sign message
	hash := sha256.Sum256(data)
	sig, err := ecc.SignBytes(privateKey, hash[:], ecc.Normal)
	if err != nil {
		panic("Signature Message Error...")
	}
	return sig
}

/**
 * 公钥验签
 * @param data  签名数据
 * @param publicKey 验签公钥
 * @param sign 签名数据
 * @return 签名是否正确
 */

func Verify(data []byte, sign []byte, publicKey *ecdsa.PublicKey) bool {
	return ecc.VerifyBytes(publicKey, data[:], sign, ecc.Normal)
}

/**
 * utxo数组（包含输入和输出）转化为byte数据供交易签名
 * @param inUtxos
 * @param outUtxos
 * @return
 */

func UTXO2Bytes(inUTXO []UTXO, outUTXO []UTXO) []byte {
	return []byte(fmt.Sprint(inUTXO) + fmt.Sprint(outUTXO))
}

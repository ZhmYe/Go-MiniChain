package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
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
 * RIPEMD160哈希摘要算法
 * @param data
 * @return
 */

func Ripemd160Digest(data []byte) []byte {
	//ripemd160 := crypto.RIPEMD160.New()
	//ripemd160.Write(data)
	//hashBytes := ripemd160.Sum(nil)
	//return hashBytes
	return data
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
	//hash := sha256.Sum256(data)
	sig, err := ecc.SignBytes(privateKey, data[:], ecc.Normal)
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

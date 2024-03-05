package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func CalculateSha256(input string) []byte {
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		panic("Failed to write input") // 处理错误情况
	}
	return hash.Sum(nil)
}

func GetSha256Digest(data string) string {
	return strings.ToUpper(hex.EncodeToString(CalculateSha256(data)))
}
func Sha256Digest(data []byte) []byte {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		panic("Failed to write input") // 处理错误情况
	}
	return hash.Sum(nil)
}

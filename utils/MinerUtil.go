package utils

import (
	"Go-Minichain/config"
	"math/rand"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 所有可能的字符集合
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
func HashPrefixTarget() string {
	result := ""
	for i := 0; i < config.MiniChainConfig.GetDifficulty(); i++ {
		result += "0"
	}
	return result
}

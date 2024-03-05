package utils

import (
	"bytes"
	"math/big"
)

type Base58Util struct {
	ALPHABET []byte
}

func NewBase58Util() *Base58Util {
	Alphabet := []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	return &Base58Util{
		ALPHABET: Alphabet,
	}
}
func (b *Base58Util) Encode(input []byte) string {
	//1. 转换成ascii码对应的值
	strByte := []byte(input)
	//2. 转换十进制
	strTen := big.NewInt(0).SetBytes(strByte)
	//fmt.Println(strTen)  // 结果4612462
	//3. 取出余数
	var modSlice []byte
	for strTen.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0) //余数
		strTen58 := big.NewInt(58)
		strTen.DivMod(strTen, strTen58, mod)                 //取余运算
		modSlice = append(modSlice, b.ALPHABET[mod.Int64()]) //存储余数,并将对应值放入其中
	}
	//  处理0就是1的情况 0使用字节'1'代替
	for _, elem := range strByte {
		if elem != 0 {
			break
		} else if elem == 0 {
			modSlice = append(modSlice, byte('1'))
		}
	}
	ReverseModSlice := b.ReverseByteArr(modSlice)
	return string(ReverseModSlice)
}
func (b *Base58Util) ReverseByteArr(bytes []byte) []byte { //将字节的数组反转
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i] //前后交换
	}
	return bytes
}
func (b *Base58Util) Decode(input string) []byte {
	strByte := []byte(input)
	//fmt.Println(strByte)  //[81 101 56 68]
	ret := big.NewInt(0)
	for _, byteElem := range strByte {
		index := bytes.IndexByte(b.ALPHABET, byteElem) //获取base58对应数组的下标
		ret.Mul(ret, big.NewInt(58))                   //相乘回去
		ret.Add(ret, big.NewInt(int64(index)))         //相加
	}
	return ret.Bytes()
}

//func (b *Base58Util) divmod(number []byte, firstDigest int, base int, divisor int) byte {
//	reminder := 0
//	for i := firstDigest; i < len(number); i++ {
//		digit := int(number[i] & 0xFF)
//		tmp := reminder*base + digit
//		number[i] = byte(tmp / divisor)
//		reminder = tmp % divisor
//	}
//	return byte(reminder)
//}

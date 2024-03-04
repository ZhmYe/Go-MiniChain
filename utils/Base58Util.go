package utils

type Base58Util struct {
	ALPHABET     string
	ENCODED_ZERO byte
	INDEXES      []int
}

func NewBase58Util() *Base58Util {
	Alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	indexes := make([]int, len(Alphabet))
	for i, char := range Alphabet {
		indexes[char] = i
	}
	return &Base58Util{
		ALPHABET:     Alphabet,
		ENCODED_ZERO: Alphabet[0],
		INDEXES:      nil,
	}
}
func (b *Base58Util) Encode(input []byte) string {
	if len(input) == 0 {
		return ""
	}
	// 统计前导0
	nbZero := 0
	for {
		if nbZero < len(input) && input[nbZero] == 0 {
			nbZero++
		} else {
			break
		}
	}
	// 最大编码数据长度
	encoded := make([]byte, len(input)*2)
	outputStart := len(encoded)
	// Base58编码正式开始
	inputStart := nbZero
	for {
		if inputStart >= len(input) {
			break
		}
		outputStart--
		encoded[outputStart] = b.ALPHABET[b.divmod(input, inputStart, 256, 58)]
		if (input[inputStart]) == 0 {
			inputStart++
		}
	}
	// 输出结果中有0,去掉输出结果的前端0
	for {
		if outputStart < len(encoded) && encoded[outputStart] == b.ENCODED_ZERO {
			outputStart++
		} else {
			break
		}
	}
	for {
		if nbZero > 0 {
			nbZero--
			outputStart--
			encoded[outputStart] = b.ENCODED_ZERO
		} else {
			break
		}
	}
	return string(encoded[outputStart:])
}
func (b *Base58Util) Decode(input string) []byte {
	if len(input) == 0 {
		return make([]byte, 0)
	}
	// 将BASE58编码的ASCII字符转换为BASE58字节序列
	input58 := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		char := input[i]
		input58[i] = char
	}
	// 统计前导0
	nbZero := 0
	for {
		if nbZero < len(input58) && input58[nbZero] == 0 {
			nbZero++
		} else {
			break
		}
	}
	// Base58 编码转 字节序（256进制）编码
	decoded := make([]byte, len(input))
	outputStart := len(decoded)
	inputStart := nbZero
	for {
		if inputStart >= len(input58) {
			break
		}
		outputStart--
		decoded[outputStart] = b.divmod(input58, inputStart, 58, 256)
		if input58[inputStart] == 0 {
			inputStart++
		}
	}
	for {
		if outputStart < len(decoded) && decoded[outputStart] == 0 {
			outputStart++
		} else {
			break
		}
	}
	return decoded[outputStart-nbZero:]
}
func (b *Base58Util) divmod(number []byte, firstDigest int, base int, divisor int) byte {

}

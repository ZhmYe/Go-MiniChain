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
	for inputStart := nbZero; inputStart < len(input); {
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

}

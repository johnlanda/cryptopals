package cryptanalysis

func SingleCharXor(character byte, in string) []byte {

	n := len(in)
	out := make([] byte, n)

	for i := 0; i < n; i++ {
		out[i] = in[i] ^ character
	}

	return out
}

func RepeatingKeyXor(key string, in string) []byte {
	keyLen := len(key)
	inLen := len(in)
	out := make([] byte, inLen)

	for i := 0; i < inLen; i ++ {
		out[i] = in[i] ^ key[i % keyLen]
	}

	return out
}

func FixedXor(s1 string, s2 string) []byte {

	in1Len := len(s1)
	in2Len := len(s2)

	var minLenStr string
	var maxLenStr string
	if in1Len < in2Len {
		minLenStr = s1
		maxLenStr = s2
	} else {
		minLenStr = s2
		maxLenStr = s1
	}

	b := make([]byte, len(maxLenStr))

	for i := 0; i < len(minLenStr); i++ {
		b[i] = minLenStr[i] ^ maxLenStr[i]
	}

	for i := len(minLenStr); i < len(maxLenStr); i++ {
		// Nothing there so whole byte is different
		b[i] = byte(0b11111111)
	}

	return b
}

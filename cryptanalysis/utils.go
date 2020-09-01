package cryptanalysis

func Transpose(slice [][]byte) [][]byte {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl - 1; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

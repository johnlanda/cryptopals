package cryptanalysis

import (
	"math/bits"
)

func HammingDistance(in1 string, in2 string) int {
	fixedXorStr := FixedXor(in1, in2)

	return countOnes(fixedXorStr)

}

func countOnes(b []byte) int {
	n := len(b)

	total := 0
	for i := 0; i < n; i++ {
		total += bits.OnesCount8(b[i])
	}

	return total
}
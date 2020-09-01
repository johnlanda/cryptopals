package cryptanalysis

import (
	"fmt"
)

func FindLikelyXorKeySize(minKeyLen int, maxKeyLen int, text []byte) int {
	minDistInit := false
	minDistNorm := float64(0)
	minDistKeyLen := 0
	for i := minKeyLen; i < maxKeyLen+1; i++ {

		numPairs := len(text) / (i * 2)
		distAccum := float64(0)
		for j := 0; j < numPairs; j++ {
			// 0 4 8 12
			// 2 6
			b1 := text[i*2*j:(i*2*j)+i]
			b2 := text[(i*2*j)+i:(i*2*j)+(2*i)]

			distAccum += float64(HammingDistance(string(b1), string(b2)))
		}


		distanceAvg := distAccum / float64(numPairs)

		distanceNorm := distanceAvg / float64(i)

		fmt.Printf("Key size: %d\nDistance (Normalized): %f\n\n", i, distanceNorm)

		if !minDistInit {
			minDistInit = true
			minDistNorm = distanceNorm
			minDistKeyLen = i
		} else if distanceNorm < minDistNorm {
			minDistNorm = distanceNorm
			minDistKeyLen = i
		}
	}

	return minDistKeyLen
}

func MostLikelySingleCharXorKey(chunk []byte) (string, float64, []byte) {

	initMin := false
	minX2Key := ""
	minX2 := float64(0)
	var likelyRes []byte

	for i := 0; i < 128; i++ {
		res := SingleCharXor(byte(i), string(chunk))
		score := ScoreAlphabet(string(res), ChiAlphaSpace)

		//fmt.Printf("%s: %f\n", []byte{byte(i)}, score)

		if !initMin {
			initMin = true
			minX2 = score
			minX2Key = string(byte(i))
			likelyRes = res
		} else if score < minX2 {
			minX2 = score
			minX2Key = string(byte(i))
			likelyRes = res
		}
	}

	return minX2Key, minX2, likelyRes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
package main

import (
	"encoding/hex"
	"fmt"
	"local.com/cryptanalysis"
	"log"
)

func main() {

	const s1 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded, err := hex.DecodeString(s1)

	if err != nil {
		log.Fatal(err)
	}

	initMin := false
	minX2Key := ""
	minX2 := float64(0)
	likelyRes := ""
	for alpha := "a"[0]; alpha <= "z"[0]; alpha++ {
		letter := string(alpha)
		res := cryptanalysis.SingleCharXor(letter[0], string(decoded))
		score := cryptanalysis.ScoreAlphabet(string(res), cryptanalysis.ChiAlphaSpace)

		if !initMin {
			initMin = true
			minX2 = score
			minX2Key = letter
			likelyRes = string(res)
		} else if score < minX2 {
			minX2 = score
			minX2Key = letter
			likelyRes = string(res)
		}
	}

	fmt.Printf("%s\n%s: %f\n", likelyRes, minX2Key, minX2)
	
}

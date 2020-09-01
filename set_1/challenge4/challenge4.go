package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"local.com/cryptanalysis"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/Users/johnlanda/Dev/personal/cryptopals/set_1/challenge4/challenge4.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	initMinOverall := false
	minX2KeyOverall := ""
	minX2Overall := float64(0)
	var likelyResOverall []byte
	for scanner.Scan() {
		minX2Key, minX2, dec := mostLikelyForLine(scanner.Text())

		//fmt.Printf("%s\n%s: %f\n", dec, minX2Key, minX2)

		if !initMinOverall {
			initMinOverall = true
			minX2KeyOverall = minX2Key
			minX2Overall = minX2
			likelyResOverall = dec
		} else if minX2 < minX2Overall {
			minX2KeyOverall = minX2Key
			minX2Overall = minX2
			likelyResOverall = dec
		}
	}

	fmt.Printf("%s\n%s: %f\n", likelyResOverall, minX2KeyOverall, minX2Overall)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mostLikelyForLine(line string) (string, float64, []byte) {

	decoded, err := hex.DecodeString(line)

	if err != nil {
		log.Fatal(err)
	}

	initMin := false
	minX2Key := ""
	minX2 := float64(0)
	var likelyRes []byte
	for i := 1; i < 256; i++ {
		res := cryptanalysis.SingleCharXor(byte(i), string(decoded))
		score := cryptanalysis.ScoreAlphabet(string(res), cryptanalysis.ChiAlphaSpace)

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

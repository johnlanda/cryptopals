package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s1 = "1c0111001f010100061a024b53535009181c"
	const s2 = "686974207468652062756c6c277320657965"

	decoded1, err := hex.DecodeString(s1)

	if err != nil {
		log.Fatal(err)
	}

	decoded2, err2 := hex.DecodeString(s2)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("decoded1: %s\n", decoded1)
	fmt.Printf("decoded2: %s\n", decoded2)

	res := fixedXor(string(decoded1), string(decoded2))

	fmt.Printf("xor result: %x\n", string(res))

}

func fixedXor(s1 string, s2 string) []byte {

	n := len(s1)
	b := make([]byte, n)

	for i := 0; i < n; i++ {
		b[i] = s1[i] ^ s2[i]
	}

	return b
}

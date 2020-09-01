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

	minX2Key, minX2, likelyRes := cryptanalysis.MostLikelySingleCharXorKey(decoded)

	fmt.Printf("%s\n%s: %f\n", likelyRes, minX2Key, minX2)
	
}

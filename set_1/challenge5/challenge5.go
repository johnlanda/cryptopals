package main

import (
	hex2 "encoding/hex"
	"fmt"
	"local.com/cryptanalysis"
)

func main() {
	plain := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	cipher := cryptanalysis.RepeatingKeyXor("ICE", plain)

	fmt.Println("%s\n", string(cipher))

	hex := hex2.EncodeToString(cipher)

	fmt.Println("%s\n", hex)
	
}

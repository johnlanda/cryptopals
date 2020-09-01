package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	encoded := hexToB64(s)

	fmt.Printf("plain encoded to base64: %s\n", encoded)
}

func hexToB64(src string) string {

  decoded, err := hex.DecodeString(src)

  if err != nil {
  	log.Fatal(err)
  }

  return base64.StdEncoding.EncodeToString(decoded)
}

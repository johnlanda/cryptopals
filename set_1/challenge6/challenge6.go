package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"local.com/cryptanalysis"
	"os"
)

func main() {

	/*
	find the keysize with the minimum normalized hamming distance
	 */

	inputPath := "/Users/johnlanda/Dev/personal/cryptopals/set_1/challenge6/challenge6.txt"
	dat, err := ioutil.ReadFile(inputPath)
	check(err)

	ciphertext, err := base64.StdEncoding.DecodeString(string(dat))
	check(err)

	minDistKeyLen := cryptanalysis.FindLikelyXorKeySize(2, 40, ciphertext)

	fmt.Printf("Most likely key size: %d\n", minDistKeyLen)

	/*
	create blocks of length keysize from file
	*/

	fileLength := len(ciphertext)
	numFullChunks := fileLength / minDistKeyLen
	chunksMod := fileLength % minDistKeyLen

	numChunks := 0
	if chunksMod != 0 {
		numChunks = numFullChunks + 1
	} else {
		numChunks = numFullChunks
	}

	// probably could pass the file handle here to the function above
	// and not have to open it twice but w/e
	f, err := os.Open(inputPath)
	check(err)

	defer f.Close()

	chunks := make([][]byte, numChunks)

	for i := 0; i < numChunks; i++ {

		if i == numChunks - 1 {
			chunks[i] = ciphertext[i*minDistKeyLen:]
		} else {
			chunks[i] = ciphertext[i*minDistKeyLen:(i+1)*minDistKeyLen]
		}

	}

	// the last row may be incomplete so easiest just to throw it out
	chunks = chunks[:len(chunks)-1]

	/*
	transpose the chunks
	 */

	chunksT := cryptanalysis.Transpose(chunks)

	fmt.Printf("Length of chunksT: %d\n", len(chunksT))

	/*
	solve each block for xor
	 */

	key := make([]byte, minDistKeyLen)
	for i := 0; i < len(chunksT); i++ {
		minX2Key, minX2, _ := cryptanalysis.MostLikelySingleCharXorKey(chunksT[i])

		fmt.Printf("%s: %f\n", minX2Key, minX2)

		key = append(key, []byte(minX2Key)...)
	}

	fmt.Printf("Predicted key: %s\n", key)

	plain := cryptanalysis.RepeatingKeyXor(string(key), string(ciphertext))

	fmt.Printf("%s\n", plain)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

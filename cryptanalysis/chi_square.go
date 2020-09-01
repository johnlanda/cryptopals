package cryptanalysis

import (
	"math"
	"strings"
)

type AlphabetFrequency struct {
	Alphabet  string
	Frequency map[rune]float64
}

var ChiAlpha = AlphabetFrequency{
	Alphabet: "abcdefghijklmnopqrstuvwxyz",
	Frequency: map[rune]float64{
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702,
		'f': 0.02228, 'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153,
		'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
		'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
		'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150, 'y': 0.01974,
		'z': 0.00074},
}

var ChiAlphaSpace = AlphabetFrequency{
	Alphabet: "abcdefghijklmnopqrstuvwxyz ",
	Frequency: map[rune]float64{
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702,
		'f': 0.02228, 'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153,
		'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
		'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
		'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150, 'y': 0.01974,
		'z': 0.00074, ' ': 0.23200},
}

func ScoreAlphabet(data string, chi AlphabetFrequency) float64 {
	data = strings.ToLower(data)
	counts := make(map[rune]int)
	chi2 := 0.0
	total := 0

	// Get a count of all the letters and the total number of letters.
	for _, c := range data {
		if strings.Contains(chi.Alphabet, string(c)) {
			total = total + 1
			_, ok := counts[c]
			if ok {
				counts[c] = counts[c] + 1
			} else {
				counts[c] = 1
			}
		}
	}

	// Do not calculate the chi-squared value unless the string is at least 70%
	// ASCII alphabet.
	////if total < int(float64(0.7)*float64(len(data))) {
	//	return 1000.0
	//}

	// Calculate chi-squared for each letter
	for _, k := range chi.Alphabet {
		expected := float64(total) * chi.Frequency[k]
		actual := float64(counts[k])
		val := math.Pow(actual-expected, 2) / expected
		chi2 = chi2 + val
	}

	return chi2
}
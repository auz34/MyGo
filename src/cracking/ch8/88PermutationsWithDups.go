// Write a method to compute all permutations of a string whose characters are not necessarily unique.
// The list of permutations should not have duplicates
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abbd"
	printPermutations(s, "")
}

func printPermutations(input, prefix string) {
	l := len(input)
	if (l == 0) {
		fmt.Printf("Permutation: %v \n", prefix)
	}

	past := make(map[rune]bool)

	for i:=0; i<l; {
		r, w := utf8.DecodeRuneInString(input[i:])
		nextI := input[0:i] + input[i+w:]
		nextP := prefix + string(r)
		i += w

		if _, ok := past[r]; !ok {
			printPermutations(nextI, nextP)
			past[r] = true
		}
	}
}



// Write a method to compute all permutations of a string of unique characters
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abcd"
	printPermutations(s, "")
}

func printPermutations(input, prefix string) {
	l := len(input)
	if (l == 0) {
		fmt.Printf("Permutation: %v \n", prefix)
	}

	for i:=0; i<l; {
		r, w := utf8.DecodeRuneInString(input[i:])
		nextI := input[0:i] + input[i+w:]
		nextP := prefix + string(r)
		printPermutations(nextI, nextP)
		i += w
	}
}


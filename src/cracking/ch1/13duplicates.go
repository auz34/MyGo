// Design an algorithm and write code to remove the duplicate characters in a string
// without using any additional buffer.
// NOTE: One or two additional variables are fine. An extra copy of the array is not.
// FOLLOW UP
// Write the test cases for this method.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	input := "hello, world"
	fmt.Print("Input string: \n")
	fmt.Printf("%v\n", input)
	fmt.Print("Output string: \n")
	fmt.Printf("%v\n", removeDuplicates(input))
}

func removeDuplicates(s string) string {
	for i, w := 0, 0; i<len(s); i += w {
		rune, width := utf8.DecodeRuneInString(s[i:])
		for j, w2 := i+width, 0; j<len(s); j += w2 {
			rune2, width2 := utf8.DecodeRuneInString(s[j:])
			if (rune2 == rune) {
				s = s[0:j] + s[j+width2:]
				w2 = 0
			} else {
				w2 = width2
			}
		}
		w = width
	}

	return s
}


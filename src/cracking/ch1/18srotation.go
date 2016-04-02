// Assume you have a method isSubstring which checks if one word is a substring of another.
// Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call
// to isSubstring (i.e., “waterbottle” is a rotation of “erbottlewat”).
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	fmt.Printf("%v \n", isRotation(s1, s2))
}

func isRotation(s1, s2 string) bool {
	firstRune, _ := utf8.DecodeRuneInString(s1[0:])

	if len(s1) != len(s2) {
		return false
	}

	for i, w := 0, 0; i<len(s2); i += w {
		rune, width := utf8.DecodeRuneInString(s2[i:])
		if rune == firstRune {
			rotationCandidate := s2[i:] + s2[0:i-1]
			if isSubstring(s1, rotationCandidate) {
				return true
			}
		}

		w = width
	}

	return false
}

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

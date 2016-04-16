/* Write a method to sort an array of strings so that all the anagrams are next to each other.*/
package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

func main() {
	list := []string { "a", "d", "q", "b", "t", "c", "e"}

	quickSort(list)
	for _, s := range list {
		fmt.Printf("%v\n", s)
	}
}

func getStringMap(s string) map[rune]int {
	result := make(map[rune]int)
	for i:=0; i<len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		result[r]++
		i+=w
	}
	return result
}

func anagramCompare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	set1 := getStringMap(s1)
	set2 := getStringMap(s2)
	for r, n := range set1 {
		if set2[r] != n {
			return false
		}
	}

	return true
}

func quickSort(list []string) {
	l := len(list)
	if l <= 1 {
		return
	}

	rand.Seed(42)
	ndx := rand.Intn(l)
	list[0], list[ndx] = list[ndx], list[0]

	sP := 1
	fP := l
	for sP < fP {
		if list[sP] < list[0] {
			sP++
		} else {
			fP--
			list[sP], list[fP] = list[fP], list[sP]
		}
	}

	if sP == l {
		quickSort(list[1:])
		return
	}

	partBorder := sP - 1
	list[0], list[partBorder] = list[partBorder], list[0]
	if ndx > 0 {
		quickSort(list[0:partBorder])
	}

	if ndx < l - 1 {
		quickSort(list[partBorder+1:])
	}
}

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

func getStringMap(s string) (map[rune]int, int) {
	result := make(map[rune]int)
	sum := 0

	for i:=0; i<len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		result[r]++
		sum += int(r)
		i+=w
	}
	return result, sum
}

func anagramCompare(s1, s2 string) int {
	if len(s1) != len(s2) {
		return false
	}

	set1, checkSum1 := getStringMap(s1)
	set2, checkSum2 := getStringMap(s2)
	if (checkSum1 != checkSum2) {
		return checkSum1 - checkSum2
	}

	for r, n := range set1 {
		if set2[r] != n {
			return -1
		}
	}

	return 0
}

func quickSort(list []string) {
	l := len(list)
	if l <= 1 {
		return
	}

	// randomizing
	rand.Seed(42)
	ndx := rand.Intn(l)
	list[0], list[ndx] = list[ndx], list[0]

	v := list[0]
	i, lt, gt := 0, 0, l-1

	for i <= gt {
		cmp := anagramCompare(v, list[i])
		if cmp<0 {
			list[lt], list[i] = list[i], list[lt]
			lt++
			i++
		} else if cmp>0 {
			list[gt], list[i] = list[i], list[gt]
			gt--
		} else {
			i++
		}
	}

	if lt>0 {
		quickSort(list[0:lt])
	}

	if gt<l-1 {
		quickSort(list[gt+1:])
	}
}

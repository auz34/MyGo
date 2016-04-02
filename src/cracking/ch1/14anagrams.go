// Write a method to decide if two strings are anagrams or not.
package main

import "fmt"

func main() {
	fmt.Printf("%v\n", isAnagram("hello", "lleho"))
	fmt.Printf("%v\n", isAnagram("hello", "leho"))
	fmt.Printf("%v\n", isAnagram("hello", "lheho"))
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}

	m2 := make(map[rune]int)
	for _, c := range s2 {
		m2[c]++
	}

	for r, count := range m1 {
		if m2[r] != count {
			return false
		}
	}

	return true
}
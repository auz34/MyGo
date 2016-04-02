// Implement an algorithm to determine if a string has all unique characters.
// What if you can not use additional data structures?
package main

import "fmt"

func main() {
	fmt.Printf("%t\n",areAllUnique("hello, world"))
	fmt.Printf("%t\n",areAllUnique("he, world"))
}

func areAllUnique(s string) bool {

	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
		if m[c] > 1 {
			return false
		}
	}

	return true
}
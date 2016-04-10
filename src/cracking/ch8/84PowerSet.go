// Write a method to return all subsets of a set

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	set := "abcd"
	printSubsets("", set)
}

func printSubsets(curSet, leftItems string) {
	fmt.Printf("Subset: {%v}\n", curSet)
	for i:=0; i < len(leftItems); {
		r, size := utf8.DecodeRuneInString(leftItems[i:])
		printSubsets(curSet+string(r), leftItems[i+size:])
		i += size
	}
}

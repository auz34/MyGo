/*
Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.

For example,
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].

Given word1 = “coding”, word2 = “practice”, return 3.
Given word1 = "makes", word2 = "coding", return 1.

Note:
You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.
*/

package main

import "fmt"

func main() {
	words := []string {"practice", "makes", "perfect", "coding", "makes"}
	fmt.Printf("%v\n", shortestDistance(words, "makes", "coding"))
	fmt.Printf("%v\n", shortestDistance(words, "coding", "practice"))
}

func findWord(words []string, word string) []int {
	result := make([]int, 0)
	for i:=0; i<len(words); i++ {
		if words[i] == word {
			result = append(result, i)
		}
	}

	return result
}

func dist(x,y int) int {
	if x>y {
		return x-y
	} else {
		return y-x
	}
}

func findMinDistance(ind1, ind2 []int) int {
	d := dist(ind1[0], ind2[0])
	for i:=0; i<len(ind1); i++ {
		for j:=0; j<len(ind2); j++ {
			if dist(ind1[i], ind2[j])<d {
				d = dist(ind1[i], ind2[j])
			}
		}
	}

	return d
}

func shortestDistance(words []string, word1 string, word2 string) int {
	ind1 := findWord(words, word1)
	ind2 := findWord(words, word2)
	return findMinDistance(ind1, ind2)
}
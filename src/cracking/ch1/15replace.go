// Write a method to replace all spaces in a string with ‘%20‘
package main

import "fmt"

func main() {
	fmt.Print(replace("hello, world"))
}

func charCount(chars []rune, char rune) int {
	cnt := 0
	for _, c := range chars {
		if c == char {
			cnt++
		}
	}

	return cnt
}

func replace(s string) string {
	chrs := []rune(s)
	originalLen := len(chrs)
	spacesCount := charCount(chrs, ' ')

	if spacesCount == 0 {
		return s
	}

	delta := spacesCount * 2
	empty := make([]rune, delta)
	chrs = append(chrs, empty...)

	for i := originalLen - 1; i>=0; i-- {
		if chrs[i] == ' ' {
			chrs[i + delta] = '0'
			chrs[i + delta - 1] = '2'
			chrs[i + delta - 2] = '%'
			delta -= 2
		} else {
			chrs[i + delta] = chrs[i]
		}

		//chrs[i] = 0

		if delta == 0 {
			break
		}
	}


	return string(chrs)
}

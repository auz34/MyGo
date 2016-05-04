/*John has discovered various rocks. Each rock is composed of various elements, and each element is represented
by a lower-case Latin letter from 'a' to 'z'. An element can be present multiple times in a rock. An element
is called a gem-element if it occurs at least once in each of the rocks.

Given the list of N rocks with their compositions, display the number of gem-elements that exist in those rocks.

Input Format
The first line consists of an integer, N, the number of rocks.
Each of the next N lines contains a rock's composition. Each composition consists of lower-case letters of
English alphabet.

Constraints
1≤N≤100

Each composition consists of only lower-case Latin letters ('a'-'z').
1≤ length of each composition ≤100

Output Format
Print the number of gem-elements that are common in these rocks. If there are none, print 0.

Sample Input
3
abcdde
baccd
eeabg

Sample Output
2

Explanation
Only "a" and "b" are the two kinds of gem-elements, since these are the only characters that occur in every rock's
composition.*/
package main

import (
	"fmt"
	"io"
	"bufio"
	"bytes"
	"unicode/utf8"
	"os"
)

func readLines(reader io.Reader, n int) []string {
	buf := bufio.NewReader(reader)
	result := make([]string, n)
	line, err := buf.ReadBytes('\n')
	curLine := 0
	for err == nil {
		line = bytes.TrimRight(line, "\n")
		if len(line) > 0 {
			if line[len(line)-1] == 13 { //'\r'
				line = bytes.TrimRight(line, "\r")
			}

			result[curLine] = string(line)
			curLine++
		}

		line, err = buf.ReadBytes('\n')
	}

	if len(line) > 0 {
		result[curLine] = string(line)
	}

	return result
}

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)

	lines := readLines(os.Stdin, n)

	fmt.Printf("%v", solve4(lines))
}

func makeMapOfRunes(lines []string) map[rune]map[int]bool {
	result := make(map[rune]map[int]bool)
	for i:=0; i<len(lines); i++ {
		line := lines[i]
		for j:=0; j<len(line); {
			r, w := utf8.DecodeRuneInString(line[j:])

			submp, ok := result[r]
			if !ok {
				submp = make(map[int]bool)
				result[r] = submp
			}

			submp[i] = true
			j += w
		}
	}

	return result
}

func solve4(lines []string) int {
	mp := makeMapOfRunes(lines)

	cnt := 0
	for _, submap := range mp {
		allTrue := true
		for i:=0; i<len(lines); i++ {
			allTrue = allTrue && submap[i]
		}

		if allTrue {
			cnt++
		}
	}

	return cnt
}
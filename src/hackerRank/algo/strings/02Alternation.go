/*Shashank likes strings in which consecutive characters are different. For example, he likes ABABA, while he
doesn't like ABAA. Given a string containing characters AA and BB only, he wants to change it into a string he likes.
To do this, he is allowed to delete the characters in the string.

Your task is to find the minimum number of required deletions.

Input Format
The first line contains an integer T, i.e. the number of test cases.
The next T lines contain a string each.

Output Format
For each test case, print the minimum number of deletions required.

Constraints
1≤T≤10
1  ≤ length of string ≤ 105
Sample Input
5
AAAA
BBBBB
ABABABAB
BABABA
AAABBB
Sample Output
3
4
0
0
4

Explanation

AAAA -> A, 3 deletions
BBBBB -> B, 4 deletions
ABABABAB -> ABABABAB, 0 deletions
BABABA -> BABABA, 0 deletions
AAABBB -> AB, 4 deletions because to convert it to AB we need to delete 2 A's and 2 B's*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"bytes"
	"unicode/utf8"
)

func readLine(reader io.Reader, f func(string)) {
	buf := bufio.NewReader(reader)
	line, err := buf.ReadBytes('\n')
	for err == nil {
		line = bytes.TrimRight(line, "\n")
		if len(line) > 0 {
			if line[len(line)-1] == 13 { //'\r'
				line = bytes.TrimRight(line, "\r")
			}
			f(string(line))
		}
		line, err = buf.ReadBytes('\n')
	}

	if len(line) > 0 {
		f(string(line))
	}
}

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i:=0; i<testCases; i++ {
		readLine(os.Stdin, func(line string) {
			fmt.Printf("%v\n", solve2(line))
		})
	}
}

func solve2(s string) int {
	prev := ' '
	cnt := 0
	for i:=0; i<len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		if prev == r {
			cnt++
		}

		i += w
		prev = r
	}

	return cnt
}
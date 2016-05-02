/*
James found a love letter his friend Harry has written for his girlfriend. James is a prankster, so he decides
to meddle with the letter. He changes all the words in the letter into palindromes.

To do this, he follows two rules:
1) He can reduce the value of a letter, e.g. he can change d to c, but he cannot change c to d.
2) In order to form a palindrome, if he has to repeatedly reduce the value of a letter, he can do it until the
letter becomes a. Once a letter has been changed to a, it can no longer be changed.

Each reduction in the value of any letter is counted as a single operation. Find the minimum number of operations
required to convert a given string into a palindrome.

Input Format
The first line contains an integer T, i.e., the number of test cases.
The next T lines will contain a string each. The strings do not contain any spaces.

Constraints
1≤T≤10
1≤ length of string ≤10^4
All characters are lower case English letters.

Output Format
A single line containing the number of minimum operations corresponding to each test case.

Sample Input
4
abc
abcba
abcd
cba
Sample Output
2
0
4
2
Explanation

For the first test case, abc -> abb -> aba.
For the second test case, abcba is already a palindromic string.
For the third test case, abcd -> abcc -> abcb -> abca = abca -> abba.
For the fourth test case, cba -> bba -> aba.
*/

package main

import (
	"fmt"
	"io"
	"bufio"
	"bytes"
	"os"
)

func readLine3(reader io.Reader, f func(string)) {
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
		readLine3(os.Stdin, func(line string) {
			fmt.Printf("%v\n", solve3(line))
		})
	}
}

func absDiff(b1, b2 byte) int {
	if b1 > b2 {
		return int(b1 - b2)
	}

	return int(b2 - b1)
}

func solve3(s string) int {
	b := []byte(s)
	res := 0

	mid := len(b) / 2
	for i:=0; i<mid; i++ {
		res += absDiff(b[i], b[len(b) - 1 - i])
	}

	return res
}
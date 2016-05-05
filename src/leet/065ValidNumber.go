/*
Validate if a given string is numeric.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one.
*/
package main

import (
	"fmt"
	"unicode/utf8"
	"unicode"
)

func main() {
	fmt.Printf("%v\n", isNumber("0"))
	fmt.Printf("%v\n", isNumber(" 0.1 "))
	fmt.Printf("%v\n", isNumber("abc"))
	fmt.Printf("%v\n", isNumber("1 a"))
	fmt.Printf("%v\n", isNumber("2e10"))
}

func isNumber(s string) bool {
	allowPoint := true
	allowE := true
	for i:=0; i<len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		i += w
		if unicode.IsDigit(r) || r == ' ' {
			continue
		}

		if r == '.' {
			if allowPoint && i < (len(s) - 1) {
				allowPoint = false
			} else {
				return false
			}
			continue
		}

		if r == 'e' {
			if allowE && i < (len(s) - 1) {
				allowE = false
			} else {
				return false
			}

			continue
		}

		return false
	}

	return true
}

// Given a boolean expression consisting of the symbols 0(false), 1(true), &(AND), | (OR), and ^(XOR), and a
// desired boolean result value "result", implement a function to count the number of parenthesizing the expression
// such it evaluates to result
// EXAMPLE
// countEval("1^0|0|1", false) -> 2
// countEval("0&0&0&1^1|1", true) -> 10

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("countEval('1^0|0|1', false) -> %v \n", countParens("1^0|0|1", false))
	fmt.Printf("countEval('0&0&0&1^1|1', true) -> %v \n", countParens("0&0&0&1^1|1", true))
}

func countParens(expr string, goal bool) int {
	if len(expr) <= 3 {
		value := false
		if expr == "1" {
			value = true
		} else if expr == "0" {
			value = false
		} else {
			value = evalExpr(expr)
		}

		if value == goal {
			return 1
		}

		return 0
	}

	if goal {
		return countTrue(expr)
	}

	return countFalse(expr)
}

func evalExpr(expr string) bool {
	r1, _ := utf8.DecodeRuneInString(expr[0:])
	op, _ := utf8.DecodeRuneInString(expr[1:])
	r2, _ := utf8.DecodeRuneInString(expr[2:])

	b1 := r1 == '1'
	b2 := r2 == '1'

	switch op {
	case '&':
		return b1 && b2
	case '|':
		return b1 || b2
	case '^':
		return (b1 || b2) && !(b1 && b2)
	}

	return false
}

func countTrue(expr string) int {
	sum := 0
	//fmt.Printf("Evaluate true for %v \n", expr)
	for i:=1; i<len(expr); {
		r, w := utf8.DecodeRuneInString(expr[i:])
		left, right := expr[0:i], expr[i+1:]
		i += w
		switch r {
		case '&':
			sum += countParens(left, true) * countParens(right, true)
		case '|':
			sum += countParens(left, true) * countParens(right, false) +
				   countParens(left, false) * countParens(right, true) +
				   countParens(left, true) * countParens(right, true)
		case '^':
			sum += countParens(left, true) * countParens(right, false) +
			countParens(left, false) * countParens(right, true)
		}
	}

	fmt.Printf("Evaluate true for %v gave %v \n", expr, sum)

	return sum
}

func countFalse(expr string) int {
	sum := 0

	for i:=1; i<len(expr); {
		r, w := utf8.DecodeRuneInString(expr[i:])
		left, right := expr[0:i], expr[i+1:]
		i += w
		switch r {
		case '&':
			sum += countParens(left, false) * countParens(right, false) +
				   countParens(left, true) * countParens(right, false) +
				   countParens(left, false) * countParens(right, true)
		case '|':
			sum += countParens(left, false) * countParens(right, false)
		case '^':
			sum += countParens(left, false) * countParens(right, false) +
			countParens(left, true) * countParens(right, true)
		}
	}

	fmt.Printf("Evaluate false for %v gave %v \n", expr, sum)

	return sum
}
// Given a boolean expression consisting of the symbols 0(false), 1(true), &(AND), | (OR), and ^(XOR), and a
// desired boolean result value "result", implement a function to count the number of parenthesizing the expression
// such it evaluates to result
// EXAMPLE
// countEval("1^0|0|1", false) -> 2
// countEval("0$0$0$1^1|1", true) -> 10

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("((1^0)|0)|1 = %v", eval("((1^0)|0)|1"))
}

func translateToBool(r rune) bool {
	if r == "0" {
		return false
	} else {
		return true
	}
}

func evalSimpleExpr(expr string) bool {
	// first rune is always 1 or 0
	r, _ := utf8.DecodeRuneInString(expr[0:])
	result := translateToBool(r)

	for i:=1; i<len(expr); i+=2 {
		opr, _ := utf8.DecodeRuneInString(expr[i:])
		r2, _ := utf8.DecodeRuneInString(expr[i+1:])
		switch opr {
		case "&":
			result = result && translateToBool(r2)
		case "|":
			result = result || translateToBool(r2)
		case "^":
			result = result ^ translateToBool(r2)
		}
	}

	return result
}

func eval(expr string) bool {
	return evalSimpleExpr(expr)
}

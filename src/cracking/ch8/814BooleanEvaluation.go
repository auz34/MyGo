// Given a boolean expression consisting of the symbols 0(false), 1(true), &(AND), | (OR), and ^(XOR), and a
// desired boolean result value "result", implement a function to count the number of parenthesizing the expression
// such it evaluates to result
// EXAMPLE
// countEval("1^0|0|1", false) -> 2
// countEval("0$0$0$1^1|1", true) -> 10

package main

import (
	"fmt"
	"cracking/ch8/basics"
)

func main() {
	infix := "((1^0)|0)|1"
	fmt.Printf("Infix expression \"%v\" can presented as \"%v\" in postfix", infix, basics.InfixToPostfix(infix))
}


/*Implement an algorithm to print all valid (e.g., properly opened and closed) combinations of n-pairs of parentheses.
EXAMPLE:
input: 3 (e.g., 3 pairs of parentheses)
output: ()()(), ()(()), (())(), ((()))*/

package main

import "fmt"

func main() {
	fmt.Print("For 3 pairs following solutions are possible: \n")
	printParens(3, 3, "")
}

func printParens(l, r int, prefix string) {
	if l == 0 && r == 0 {
		fmt.Printf("Solution: %v \n", prefix)
		return
	}

	if l==0 {
		printParens(l, r-1, prefix + ")")
		return
	}

	if (l < r) {
		printParens(l-1, r, prefix + "(")
		printParens(l, r-1, prefix + ")")
		return
	}

	printParens(l-1, r, prefix + "(")
}
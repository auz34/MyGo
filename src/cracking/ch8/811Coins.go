// Given an infinite number of quarters (25 cents), dimes (10 cents), nickels (5 cents) and pennies (1 cent),
// write code to calculate the number of ways of representing n cents.
package main

import "fmt"

func main() {
	fmt.Printf("There are %v ways to change 3 cents\n", changeWays(3, 25))
	fmt.Printf("There are %v ways to change 5 cents\n", changeWays(5, 25))
	fmt.Printf("There are %v ways to change 10 cents\n", changeWays(10, 25))
	fmt.Printf("There are %v ways to change 25 cents\n", changeWays(25, 25))
	fmt.Printf("There are %v ways to change 100 cents\n", changeWays(100, 25))
}

func changeWays(amount, denom int) int {
	if amount == 0 {
		return 1
	}

	next_denom := 0

	switch denom {
	case 25: next_denom = 10
	case 10: next_denom = 5
	case 5: next_denom = 1
	case 1: return 1
	}

	sum := changeWays(amount, next_denom)
	for i:=1; i*denom<=amount; i++ {
		sum += changeWays(amount - i*denom, next_denom)
	}

	return sum
}
/*You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a
method to set all bits between i and j in N equal to M (e.g., M becomes a substring of
N located at i and starting at j).
EXAMPLE:
Input: N = 10000000000, M = 10101, i = 2, j = 6
Output: N = 10001010100*/

package main

import "fmt"

func main() {
	n := uint(1024)
	m := uint(21)
	fmt.Printf("Result %b \n", copyInBetween(n, m, uint(2), uint(6)))
}

func copyInBetween(N, M, left, right uint) uint {
	mask := ^uint(0)
	fmt.Printf("Mask: %b \n", mask)
	shift := M << left
	fmt.Printf("Shifted: %b (%v) \n", shift, shift)
	nshift := ^shift
	fmt.Printf("n-Shifted: %b (%v) \n", nshift, nshift)
	return N | shift
}

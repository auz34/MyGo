// Write a program to swap odd and even bits in an integer with as few instructions as possible
// (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, etc).

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Swap for 5 is %v", swap(5))
}

func swap(n uint) uint {
	oddmask64, _ := strconv.ParseUint("01010101010101010101010101010101", 2, 32)
	evenmask64, _ := strconv.ParseUint("10101010101010101010101010101010", 2, 32)

	odds := n & uint(oddmask64)
	evens:= n & uint(evenmask64)

	return (odds << 1) | (evens >> 1)
}
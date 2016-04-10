// An array A[1...n] contains all the integers from 0 to n except for one number which is missing.
// In this problem, we cannot access an entire integer in A with a single operation. The elements of A are
// represented in binary, and the only operation we can use to access them is “fetch the jth bit of A[i]”,
// which takes constant time. Write code to  nd the missing integer. Can you do it in O(n) time?

package main

import (
	"fmt"
)

func main() {
	array := []uint {0, 1, 2, 3, 5, 6, 7}
	fmt.Printf("Missed item = %v \n", findMissing(array))
}

func getBit(ar []uint, itemIndex int, bitIndex uint) bool {
	item := ar[itemIndex]
	return (item & (uint(1) << bitIndex)) != 0
}

func findMissing(ar []uint) uint {
	N := len(ar)
	sum, expectedSum := uint(0), uint(N)

	for i:=0; i<N; i++ {
		expectedSum += uint(i)
		for j:=uint(0); j<32; j++ {
			if getBit(ar, i, j) {
				sum += uint(1) << j
			}
		}
	}

	return expectedSum - sum
}

// Write a function to determine the number of bits required to convert integer A to integer B.
// Input: 31, 14 Output: 2
package main

import "fmt"

func main() {
	fmt.Printf("Distance between %v and %v is %v \n", 31, 14, distance(31, 14))
}

func distance(i1, i2 uint) int {
	diff := i1 &^ i2

	cnt := 0
	for i:=uint(0); i<32; i++ {
		if (diff & (uint(1) << i)) != 0 {
			cnt++
		}
	}

	return cnt
}
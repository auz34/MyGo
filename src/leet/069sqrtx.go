/*
Implement int sqrt(int x).

Compute and return the square root of x.
 */

package main

import "fmt"

func main() {
	//fmt.Printf("sgrt(9) = %v\n", mySqrt(9))
	fmt.Printf("sgrt(32768) = %v\n", mySqrt(32768))
	//fmt.Printf("sgrt(16384) = %v\n", mySqrt(16384))
	fmt.Printf("sgrt(2) = %v\n", mySqrt(2))
}

func binarySqrt(x int, left, right int) int {
	if right < left {
		return right
	}

	mid := (left + right) / 2

	//fmt.Printf("left:%v; right:%v; mid:%v \n", left, right, mid)

	midSqrt := mid*mid
	if (midSqrt == x) {
		return mid
	}

	if midSqrt < x {
		return binarySqrt(x, mid + 1, right)
	}

	return binarySqrt(x, left, mid-1)
}

func mySqrt(x int) int {
	if (x < 0) {
		return -1
	}

	right := int(1<<16)
	return binarySqrt(x, 0, right)
}
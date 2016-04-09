// Given an integer, print the next smallest and next largest number that have
// the same number of 1 bits in their binary representation.
package main

import "fmt"

func main() {

	a2, _ := nextLargest(uint(2))
	fmt.Printf("Next largest for 2 is %v \n", a2)
	a5, _ := nextLargest(uint(5))
	fmt.Printf("Next largest for 5 is %v \n", a5)
	a7, _ := nextLargest(uint(7))
	fmt.Printf("Next largest for 7 is %v (%b) \n", a7, a7)

	b4, _ := nextSmallest(uint(4))
	fmt.Printf("Next smallest for 4 is %v \n", b4)
	b6, _ := nextSmallest(uint(6))
	fmt.Printf("Next smallest for 6 is %v \n", b6)
	b14, _ := nextSmallest(uint(14))
	fmt.Printf("Next smallest for 14 is %v (%b) \n", b14, b14)
}

const one = uint(1)

func isOne(source, index uint) bool {
	return (one << index) & source != 0
}

func setOne(source *uint, index uint) {
	*source = *source | (one << index)
}

func setZero(source *uint, index uint) {
	*source = *source & ^(one << index)
}

func nextLargest(source uint) (uint, bool) {
	i:=uint(0)
	smallestOne := -1
	for i<32 && smallestOne == -1 {
		if isOne(source, i) {
			smallestOne = int(i)
		}

		i++
	}

	if smallestOne == -1 {
		return uint(0), false
	}

	smallestZeroWithOneBehind := -1
	for i<32 && smallestZeroWithOneBehind == -1 {
		if !isOne(source, i) {
			smallestZeroWithOneBehind = int(i)
		}

		i++
	}

	if smallestZeroWithOneBehind == -1 {
		return uint(0), false
	}

	oneBehind := -1
	for i >= 0 && oneBehind == -1 {
		if isOne(source, i) {
			oneBehind = int(i)
		}

		i--
	}

	result := source
	setZero(&result, uint(oneBehind))
	setOne(&result, uint(smallestZeroWithOneBehind))

	return result, true
}

func nextSmallest(source uint) (uint, bool) {
	i:=uint(0)
	smallestZero := -1
	for i<32 && smallestZero == -1 {
		if !isOne(source, i) {
			smallestZero = int(i)
		}

		i++
	}

	if smallestZero == -1 {
		return uint(0), false
	}

	smallestOneWithZeroBehind := -1
	for i<32 && smallestOneWithZeroBehind == -1 {
		if isOne(source, i) {
			smallestOneWithZeroBehind = int(i)
		}

		i++
	}

	if smallestOneWithZeroBehind == -1 {
		return uint(0), false
	}

	zeroBehind := -1
	for i>=0 && zeroBehind == -1 {
		if !isOne(source, i) {
			zeroBehind = int(i)
		}
		i--
	}

	result := source
	setZero(&result, uint(smallestOneWithZeroBehind))
	setOne(&result, uint(zeroBehind))

	return result, true
}
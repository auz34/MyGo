// A magic index in an array A[0...n-1] is defined to be an index such that A[i] = i. Given a sorted array
// of distinct integers, write a method to find a magic index, if one exists, in array A.
// FOLLOW UP
// What if the values are not distinct?
package main

import "fmt"

func main() {
	ar1 := []int {-10, -2, 1, 3, 6, 12, 15, 17, 18 }
	ar2 := []int {-10, -2, -1, 1, 3, 4, 5, 7, 18 }
	ar3 := []int {-10, 2, 2, 2, 2, 4, 5, 9, 18 }


	fmt.Printf("Brute force answers: %v, %v, %v \n", bruteForce(ar1), bruteForce(ar2), bruteForce(ar3))

	fmt.Printf("Answer for first array: %v \n", findDistinct(ar1, 0, len(ar1) - 1))
	fmt.Printf("Answer for second array: %v \n", findDistinct(ar2, 0, len(ar2) - 1))

	fmt.Printf("Answer for third array: %v \n", findNonDistinct(ar3, 0, len(ar3) - 1))
}

func bruteForce(array []int) int {
	for i:=0; i<len(array); i++ {
		if array[i]==i {
			return i
		}
	}

	return -1
}

func findDistinct(array []int, left, right int) int {
	if left>right {
		return -1
	}

	mid := (left + right) / 2

	if mid == array[mid] {
		return mid
	}

	if mid < array[mid] {
		return findDistinct(array, left, mid - 1)
	}

	return findDistinct(array, mid+1, right)
}

func findNonDistinct(array []int, left, right int) int {
	if left>right {
		return -1
	}

	mid := (left + right) / 2

	if mid == array[mid] {
		return mid
	}

	if mid < array[mid] {
		diff := array[mid] - mid
		lr := findNonDistinct(array, left, mid - 1)
		if lr == -1 {
			return findNonDistinct(array, mid + diff, right)
		} else {
			return lr
		}
	} else {
		diff := mid - array[mid]
		rr := findNonDistinct(array, mid+1, right)
		if rr == -1 {
			return findNonDistinct(array, left, mid - diff)
		} else {
			return rr
		}
	}
}
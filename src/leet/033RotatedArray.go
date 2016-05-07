/*
Suppose a sorted array is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.
*/

package main

import "fmt"

func main() {
	ar := []int {4, 5, 6, 7, 8, 0, 1, 2}

	fmt.Printf("4 -> %v\n", search(ar, 4))
	fmt.Printf("5 -> %v\n", search(ar, 5))
	fmt.Printf("6 -> %v\n", search(ar, 6))
	fmt.Printf("7 -> %v\n", search(ar, 7))
	fmt.Printf("8 -> %v\n", search(ar, 8))
	fmt.Printf("0 -> %v\n", search(ar, 0))
	fmt.Printf("1 -> %v\n", search(ar, 1))
	fmt.Printf("2 -> %v\n", search(ar, 2))

	fmt.Printf("3 -> %v\n", search(ar, 3))
	fmt.Printf("9 -> %v\n", search(ar, 9))
}

func searchInRange(nums []int, target, left, right int) int {
	if right < left {
		return -1
	}


	mid := (left + right) / 2
	comp := nums[mid] - target
	if comp<0 {
		// search right
		ndx1 := searchInRange(nums, target, mid+1, right)
		if ndx1 == -1 && nums[left]>nums[right] {
			return searchInRange(nums, target, left, mid-1)
		}

		return ndx1
	} else if comp > 0 {
		ndx1 := searchInRange(nums, target, left, mid-1)
		if ndx1 == -1 && nums[left]>nums[right] {
			return searchInRange(nums, target, mid+1, right)
		}

		return ndx1
	} else {
		return mid
	}
}

func search(nums []int, target int) int {
	return searchInRange(nums, target, 0, len(nums) - 1)
}
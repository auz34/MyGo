/*
Given a sorted array of integers, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/

package main

import "fmt"

func main() {
	ar := []int { 5, 7, 7, 8, 8, 10}
	fmt.Printf("%v\n", searchRange(ar, 8))
	fmt.Printf("%v\n", searchRange(ar, 6))
}

func searchInRange34(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	comp := nums[mid] - target
	if comp == 0 {
		return mid
	}

	if comp > 0 {
		return searchInRange34(nums, target, left, mid - 1)
	}

	return searchInRange34(nums, target, mid+1, right)
}


func searchRange(nums []int, target int) []int {
	ndx := searchInRange34(nums, target, 0, len(nums) - 1)
	if ndx == -1 {
		return []int {-1, -1}
	}

	left := ndx
	for left>=0 && nums[left] == target {
		left--
	}

	right := ndx
	for right<len(nums) && nums[right] == target {
		right++
	}

	return []int{left+1, right-1}
}
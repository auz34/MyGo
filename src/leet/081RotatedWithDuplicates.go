/*
Follow up for "Search in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?

Write a function to determine if a given target is in the array.
 */
package main

import "fmt"

func main() {
	fmt.Printf("%v", search([]int { 7, 7, 1, 1, 2, 3, 3, 3, 4, 4, 5, 6, 7 }, 2))
}

func internalSearch(nums []int, target, left, right int) bool {
	if (right < left) {
		return false
	}

	mid := (left + right) / 2

	if nums[mid] == target {
		return true
	}

	rotated := nums[left] >= nums[right]
	if nums[mid] < target {
		if rotated {
			return internalSearch(nums, target, mid+1, right) || internalSearch(nums, target, left, mid-1)
		}

		return internalSearch(nums, target, mid+1, right)
	}

	if rotated {
		return internalSearch(nums, target, left, mid-1) || internalSearch(nums, target, mid+1, right)
	}

	return internalSearch(nums, target, left, mid-1)
}

func search(nums []int, target int) bool {
	return internalSearch(nums, target, 0, len(nums) - 1)
}

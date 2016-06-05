/*
Follow up for "Find Minimum in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?
Suppose a sorted array is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Find the minimum element.

The array may contain duplicates.
 */

package main

import "fmt"

func main () {
	fmt.Printf("%v\n", findMin([]int {3, 1, 3}))
}

func min154(x, y int) int {
	if x<y {
		return x
	}

	return y
}

func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums) - 1] {
		return nums[0]
	}

	mid := len(nums) / 2

	return min154(findMin(nums[:mid]), findMin(nums[mid:]))
}
/*Suppose a sorted array is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Find the minimum element.

You may assume no duplicate exists in the array.*/

package main

import "fmt"

func main () {
	fmt.Printf("%v\n", findMin([]int {4, 5, 6, 7, 0, 1, 2}))
}

func min153(x, y int) int {
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

	return min153(findMin(nums[:mid]), findMin(nums[mid:]))
}
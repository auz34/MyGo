/*
Given an array of n integers where n > 1, nums, return an array output such that output[i] is equal to the product of
all the elements of nums except nums[i].

Solve it without division and in O(n).

For example, given [1,2,3,4], return [24,12,8,6].

Follow up:
Could you solve it with constant space complexity? (Note: The output array does not count as extra space for the
purpose of space complexity analysis.)
*/
package main

import "fmt"

func main() {
	ar := []int {1,2,3,4}

	fmt.Printf("Result: %v", productExceptSelf(ar))
}

func productExceptSelf(nums []int) []int {
	bk := make([]int, len(nums))
	result := make([]int, len(nums))

	for i:=len(nums) - 1; i>0; i-- {
		prev := 1
		if i<len(nums) - 1 {
			prev = bk[i+1]
		}

		bk[i] = prev*nums[i]
	}

	prev := 1
	for i:=0; i< len(nums); i++ {
		next := 1
		if i<len(nums) - 1 {
			next = bk[i+1]
		}

		result[i] = prev*next
		prev = prev * nums[i]
	}

	return result
}
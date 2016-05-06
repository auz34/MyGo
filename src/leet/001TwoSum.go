/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

func main() {
	ar := []int {3, 2, 4}
	fmt.Printf("%v \n", twoSum(ar, 6))
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		dict[nums[i]] = i
	}

	for i:=0; i<len(nums); i++ {
		ndx, ok := dict[target - nums[i]]
		if ok && ndx != i {
			return []int {i, ndx}
		}
	}

	return []int{}
}

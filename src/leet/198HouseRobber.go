/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and
it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of
money you can rob tonight without alerting the police.
*/

package main

import "fmt"

func main() {
	nums := []int {1, 7, 1, 1, 8}
	fmt.Printf("%v\n", rob(nums))
}

func max198(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func rob(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}

	if (len(nums) == 1) {
		return nums[0]
	}

	if (len(nums) == 2) {
		return max198(nums[0], nums[1])
	}

	mem0 := nums[0]
	mem1 := nums[1]
	mem2 := nums[2] + nums[0]
	for i:=3; i<len(nums); i++ {
		next := max198(mem0, mem1) + nums[i]
		mem0, mem1, mem2 = mem1, mem2, next
	}

	return max198(mem2, mem1)

}

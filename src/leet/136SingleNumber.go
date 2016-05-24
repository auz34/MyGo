/*
Given an array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 */
package main

import "fmt"

func main() {
	fmt.Printf("%v\n", singleNumber([]int {1, 2, 3, 2, 1, 3, 4}))
}

func singleNumber(nums []int) int {
	dict := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		dict[nums[i]]++
	}

	for k, cnt := range dict {
		if cnt == 1 {
			return k
		}
	}

	return -1
}

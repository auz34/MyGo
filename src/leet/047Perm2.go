/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
[1,1,2], [1,2,1], and [2,1,1].
 */

package main

import "fmt"

func main() {
	ar := []int {1,2,1}

	result := permuteUnique(ar)

	for i:=0; i<len(result); i++ {
		fmt.Printf("%v\n", result[i]);
	}
}


func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int { nums }
	}

	result := make([][]int, 0)
	dict := make(map[int]bool)

	for i:=0; i<len(nums); i++ {
		num := nums[i]
		if dict[num] {
			continue
		}

		dict[num] = true
		rem := make([]int, len(nums))
		copy(rem, nums)
		rem = append(rem[:i], rem[i+1:]...)

		subresult := permuteUnique(rem)
		for j:=0; j<len(subresult); j++ {
			single := append(subresult[j], nums[i])
			result = append(result, single)
		}
	}

	return result
}


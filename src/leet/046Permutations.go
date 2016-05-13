/*
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], and [3,2,1].
 */
package main

import "fmt"

func main() {
	ar := []int {1,2,1}

	result := permute(ar)

	for i:=0; i<len(result); i++ {
		fmt.Printf("%v\n", result[i]);
	}
}


func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int { nums }
	}

	result := make([][]int, 0)

	for i:=0; i<len(nums); i++ {
		rem := make([]int, len(nums))
		copy(rem, nums)
		rem = append(rem[:i], rem[i+1:]...)

		subresult := permute(rem)
		for j:=0; j<len(subresult); j++ {
			single := append(subresult[j], nums[i])
			result = append(result, single)
		}
	}

	return result
}

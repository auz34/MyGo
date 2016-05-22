/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:
Elements in a triplet (a,b,c) must be in non-descending order. (ie, a ≤ b ≤ c)
The solution set must not contain duplicate triplets.
    For example, given array S = {-1 0 1 2 -1 -4},

    A solution set is:
    (-1, 0, 1)
    (-1, -1, 2)
*/

package main

import "fmt"

func main() {
	ar := []int {-1, 0, 1, 2, -1, -4}

	res := threeSum(ar)
	for i:=0; i<len(res); i++ {
		fmt.Printf("%v\n", res[i])
	}
}

func min015(a, b int) int {
	if a<b {
		return a
	}

	return b
}

func max015(a, b int) int {
	if a>b {
		return a
	}

	return b
}

func sortThree(a, b, c int) []int {
	if a < b {
		if a<c {
			return []int { a, min015(b, c), max015(b, c)}
		}

		return []int { c, a, b }
	}

	if b<c {
		return []int { b, min015(a, c), max015(a, c)}
	}

	return []int { c, b, a }

}

func threeSum(nums []int) [][]int {
    type tupple struct {
		first int
		second int
	}

	dict := make(map[int]int)
	res := make([][]int, 0)
	for i:=0; i<len(nums); i++ {
		dict[nums[i]] = i
	}

	duplicateDetector := make(map[tupple]bool)

	for i:=0; i<len(nums) - 2; i++ {
		for j:=i+1; j<len(nums) - 1; j++ {
			target := -nums[i] - nums[j]
			item, ok := dict[target]
			if ok && item > j {
				candidate := sortThree(nums[i], nums[j], target)
				t := tupple{candidate[0], candidate[1]}
				_, dupl := duplicateDetector[t]
				if (!dupl) {
					res = append(res, candidate)
					duplicateDetector[t] = true
				}
			}
		}
	}

	return res
}
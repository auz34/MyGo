/*Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [−2,1,−3,4,−1,2,1,−5,4],
the contiguous subarray [4,−1,2,1] has the largest sum = 6*/

package main

import "fmt"

func main() {
	ar := []int { -2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("%v", maxSubArray(ar))
}

func maxSubArray(nums []int) int {
	prev := nums[0]
	result := prev
	for i:=1; i<len(nums); i++ {
		if prev > 0 {
			prev = nums[i] + prev
		} else {
			prev = nums[i]
		}

		if prev > result {
			result = prev
		}
	}

	return result
}
/*
Find the contiguous subarray within an array (containing at least one number) which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.
*/

package main

import "fmt"

func main() {
	ar:= []int {2,3,-2,4}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar = []int {-1, -2, -9, -6}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar = []int {-1, 0}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar = []int {-1, -2, -9}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar = []int {-1, -2, -9, 0, 7, 6}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar= []int {0, 0, -3, 1}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))

	ar= []int {2, -5, -2, -4, 3}
	fmt.Printf("%v -> %v\n", ar, maxProduct(ar))
}

func max152(x,y int) int {
	if x>y {
		return x
	}

	return y
}

func threeMax(x, y, z int) int {
	xy := max152(x,y)
	return max152(xy, z)
}

func abs(x int) int {
	if x<0 {
		return -x
	}

	return x
}

func absMax(x, y int) int {
	if abs(x) > abs(y) {
		return x
	}

	return y
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] == 0 {
		return max152(0, maxProduct(nums[1:]))
	}

	prev := nums[0]
	absPrev := nums[0]
	result := prev

	negNdx := -1
	negSum := absPrev
	if nums[0] < 0 {
		negNdx = 0
	}

	for i:=1; i<len(nums); i++ {
		prev = threeMax(prev*nums[i], nums[i], absPrev*nums[i])
		if negNdx != -1 {
			prev = max152(prev, (absPrev*nums[i]) / negSum)
		}

		absPrev = absMax(absPrev*nums[i], nums[i])

		if negNdx == -1 && nums[i]<0 {
			negSum = absPrev
			negNdx = i
		}


		if prev > result {
			result = prev
		}

		if nums[i] == 0 && i<(len(nums)-1) {
			return max152(result, maxProduct(nums[i+1:]))
		}
	}

	return result
}
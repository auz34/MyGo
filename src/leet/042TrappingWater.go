/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
 */

package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", trap([]int { 0,1,0,2,1,0,1,3,2,1,2,1 }))
	//fmt.Printf("%v\n", trap([]int { 4, 2, 3 }))
	fmt.Printf("%v\n", trap([]int { 5, 4, 1, 2 }))
}

func getOneTrapSum(height []int, left, right int) int {
	minHeight := height[left]
	if (height[right] < minHeight) {
		minHeight = height[right]
	}

	trapSum := 0
	for i:=left+1; i<right; i++ {
		if (minHeight > height[i]) {
			trapSum += minHeight - height[i]
		}
	}

	return trapSum
}

func getRightBorder(height []int, left, right int) int {
	if height[left] <= height[right] {
		return right
	}

	j := right+1
	for  j<len(height) {
		if height[j] > height[right] {
			return getRightBorder(height, left, j)
		}

		j++
	}

	return right
}

func trap(height []int) int {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	for i:=0; i<len(height)-2; i++ {
		if height[i] > height[i+1] {
			lefts = append(lefts, i)
		}
	}

	for i:=2; i<len(height); i++ {
		if height[i] > height[i-1] {
			rights = append(rights, i)
		}
	}

	lP, rP := 0, 0
	sum := 0
	for lP < len(lefts) && rP < len(rights) {
		lNdx, rNdx := lefts[lP], rights[rP]
		if lNdx >= rNdx {
			rP++
			continue
		}


		rNdx = getRightBorder(height, lNdx, rNdx)
		sum += getOneTrapSum(height, lNdx, rNdx)

		for lP < len(lefts) && lefts[lP] < rNdx {
			lP++
		}
	}

	return sum
}
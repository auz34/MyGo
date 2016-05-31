/*
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines
are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis
forms a container, such that the container contains the most water.

Note: You may not slant the container.
 */

package main

import "fmt"

func main() {
	fmt.Printf(" %v  \n", maxArea([]int {1, 2, 1}))
	fmt.Printf(" %v  \n", maxArea([]int {1, 4, 5}))
}

func min011(x, y int) int{
	if x < y {
		return x
	}

	return y
}

func maxArea(height []int) int {
	result := 0
	i, j := 0, len(height) - 1
	for i < j {
		curHeight := min011(height[i], height[j]) * (j - i)
		if curHeight > result {
			result = curHeight
		}

		if (height[i] < height[j]) {
			i++
		} else {
			j--
		}
	}

	return result
}

/*There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of
painting each house with a certain color is different.
You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example,
costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color
green, and so on... Find the minimum cost to paint all houses.

Note:
All costs are positive integers.*/

package  main

import "fmt"

func main() {
	costs := [][]int {
		[]int{1,1,1},
		[]int{2,1,2},
		[]int{8,1,8},
		[]int{1,1,1},
	}

	fmt.Printf("%v\n", minCost(costs))
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func minCost(costs [][]int) int {
	prev := make([]int, 3)

	for i:=0; i<len(costs); i++ {
		red := min(prev[1], prev[2]) + costs[i][0]
		blue := min(prev[0], prev[2]) + costs[i][1]
		green := min(prev[0], prev[1]) + costs[i][2]

		prev[0] = red
		prev[1] = blue
		prev[2] = green
	}

	result := min(prev[0], prev[1])
	result = min(result, prev[2])

	return result
}

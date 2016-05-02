/*You are situated in an N dimensional grid at position (x1,x2,...,xN). The dimensions of the grid are (D1,D2,...DN).
In one step, you can walk one step ahead or behind in any one of the N dimensions. (So there are always 2×N possible
different moves). In how many ways can you take M steps such that you do not leave the grid at any point? You leave
the grid if at any point xi, either xi≤0 or xi>Di.

Input Format
The first line contains the number of test cases T. T test cases follow. For each test case, the first line contains
N and M, the second line contains x1,x2,…,xN and the 3rd line contains D1,D2,…,DN.

Output Format
Output T lines, one corresponding to each test case. Since the answer can be really huge, output it modulo 1000000007.

Constraints
1≤T≤10
1≤N≤10
1≤M≤300
1≤Di≤100
1≤xi≤Di

Sample Input

1
2 3
1 1
2 3
Sample Output
12

Explanation

Starting from (1, 1) in a 2×3 2-D grid, and need to count the number of possible paths with length equal to 3. Here are
the 12 paths:

(1, 1) -> (1, 2) -> (1, 1) -> (1, 2)
(1, 1) -> (1, 2) -> (1, 1) -> (2, 1)
(1, 1) -> (1, 2) -> (1, 3) -> (1, 2)
(1, 1) -> (1, 2) -> (1, 3) -> (2, 3)
(1, 1) -> (1, 2) -> (2, 2) -> (1, 2)
(1, 1) -> (1, 2) -> (2, 2) -> (2, 1)
(1, 1) -> (1, 2) -> (2, 2) -> (2, 3)
(1, 1) -> (2, 1) -> (1, 1) -> (1, 2)
(1, 1) -> (2, 1) -> (1, 1) -> (2, 1)
(1, 1) -> (2, 1) -> (2, 2) -> (2, 1)
(1, 1) -> (2, 1) -> (2, 2) -> (1, 2)
(1, 1) -> (2, 1) -> (2, 2) -> (2, 3)*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i:=0; i<testCases; i++ {
		var n int
		fmt.Fscan(os.Stdin, &n)

		var m int
		fmt.Fscan(os.Stdin, &m)

		x := make([]int, n)
		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			x[j] = v-1
		}

		d := make([]int, n)
		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			d[j] = v-1
		}

		fmt.Printf("%v\n", solve6(m, x, d))
	}
}

func getStepMap(previousStep map[uint64]int, d []int) map[uint64]int {
	result := make(map[uint64]int)
	x := make([]int, len(d))

	ok := true
	for ok {
		key := cacheKey(x)
		sum := getPointSum(previousStep, x, d)

		result[key] = sum

		ok = next(x, d)
	}

	return result
}

func getPointSum(previousStep map[uint64]int, x, d []int) int {
	sum := 0
	for i:=0; i<len(x); i++ {
		if x[i] > 0 {
			if previousStep != nil {
				x[i]--
				key1 := cacheKey(x)
				sum += previousStep[key1]
				x[i]++
			} else {
				sum++
			}
		}

		if x[i] < d[i] {
			if previousStep != nil {
				x[i]++
				key1 := cacheKey(x)
				sum += previousStep[key1]
				x[i]--
			} else {
				sum++
			}
		}
	}

	return sum % 1000000007
}

func next(x, d []int) bool {
	curDim := 0
	ok := false
	for curDim<len(d) && !ok {
		if x[curDim] < d[curDim] {
			x[curDim]++
			ok = true
		} else {
			x[curDim] = 0
			curDim++
		}
	}

	return ok
}

func solve6(m int, x, d []int) int {
	if m == 0 {
		return 1
	}

	var previousStep map[uint64]int

	for i:=1; i<m; i++ {
		previousStep = getStepMap(previousStep, d)
	}

	return getPointSum(previousStep, x, d)
}

func cacheKey(x []int) uint64 {
	m := uint64(1)
	key := uint64(0)
	for i:=0; i<len(x); i++ {
		key += uint64(x[i]) * m
		m *= 100
	}

	return key
}
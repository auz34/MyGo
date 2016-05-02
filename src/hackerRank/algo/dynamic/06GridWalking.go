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
	"strconv"
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
			x[j] = v
		}

		d := make([]int, n)
		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			d[j] = v
		}

		cache := make(map[string]int)
		fmt.Printf("%v\n", solve6(m, x, d, cache))
	}
}

func solve6(m int, x, d []int, cache map[string]int) int {
	if m == 0 {
		return 1
	}

	key := cacheKey(m, x)
	res, ok := cache[key]
	if ok {
		return res
	}

	sum := 0
	for i:=0; i<len(x); i++ {
		if x[i] > 1 {
			x[i]--
			sum += solve6(m-1, x, d, cache)
			x[i]++
		}

		if x[i] < d[i] {
			x[i]++
			sum += solve6(m-1, x, d, cache)
			x[i]--
		}
	}

	return sum
}

func cacheKey(m int, x []int) string {
	s := strconv.Itoa(m) + "_"
	for i:=0; i<len(x); i++ {
		s += strconv.Itoa(x[i]) + "_"
	}

	return s
}
/* Given an array A={a1,a2,…,aN} of N elements, find the maximum possible sum of a
1) Contiguous subarray
2) Non-contiguous (not necessarily contiguous) subarray.

Empty subarrays/subsequences should not be considered.

Input Format
First line of the input has an integer T. T cases follow.
Each test case begins with an integer N. In the next line, N integers follow representing the elements of array A.

Constraints:
1≤T≤101≤T≤10
1≤N≤1051≤N≤105
−104≤ai≤104−104≤ai≤104
The subarray and subsequences you consider should have at least one element.

Output Format
Two, space separated, integers denoting the maximum contiguous and non-contiguous subarray.
At least one integer should be selected and put into the subarrays (this may be required in cases where all elements are negative).

Sample Input
2
4
1 2 3 4
6
2 -1 2 3 4 -5
Sample Output

10 10
10 11
Explanation

In the first case:
The max sum for both contiguous and non-contiguous elements is the sum of ALL the elements (as they are all positive).

In the second case:
[2 -1 2 3 4] --> This forms the contiguous sub-array with the maximum sum.
For the max sum of a not-necessarily-contiguous group of elements, simply add all the positive elements.*/
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
		ar := make([]int, n)
		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			ar[j] = v
		}
		res1, res2 := solve2(ar)
		fmt.Printf("%v %v\n", res1, res2)
	}
}

func max(a, b int) int {
	if a>b {
		return a
	}

	return b
}

func solve2(ar []int) (int, int) {
	mem := make([]int, len(ar))
	mem[0] = ar[0]
	res1 := mem[0]
	maxItem := mem[0]
	positiveSum := 0
	for i:=1; i<len(ar); i++ {
		mem[i] = max(mem[i-1] + ar[i], ar[i])
		if mem[i] > res1 {
			res1 = mem[i]
		}

		if ar[i] > maxItem {
			maxItem = ar[i]
		}

		if ar[i] > 0 {
			positiveSum += ar[i]
		}
	}

	if positiveSum > 0 {
		return res1, positiveSum
	}

	return res1, maxItem
}


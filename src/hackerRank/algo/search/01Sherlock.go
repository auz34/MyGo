/*
Watson gives Sherlock an array A of length N. Then he asks him to determine if there exists an element in the array
such that the sum of the elements on its left is equal to the sum of the elements on its right.
If there are no elements to the left/right, then the sum is considered to be zero.
Formally, find an ii, such that, A1++A2...A...Ai-1 =A=Ai+1+A+Ai+2...A...AN.

Input Format
The first line contains TT, the number of test cases.
For each test case, the first line contains NN, the number of elements in the array AA.
The second line for each test case contains NN space-separated integers, denoting the array AA.

Output Format
For each test case print YES if there exists an element in the array, such that the sum of the elements
on its left is equal to the sum of the elements on its right; otherwise print NO.

Constraints
1≤T≤101≤T≤10
1≤N≤1051≤N≤105
1≤A1≤Ai ≤2×104≤2×104
1≤i≤N1≤i≤N
Sample Input

2
3
1 2 3
4
1 2 3 3
Sample Output

NO
YES
Explanation
For the first test case, no such index exists.
For the second test case, A[1]+A[2]=A[4]A[1]+A[2]=A[4], therefore index 33 satisfies the given conditions.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i:=0; i<testCases; {
		var n int
		fmt.Fscan(os.Stdin, &n)
		ar := make([]int, n)
		arSum := 0

		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			ar[j] = v
			arSum += v
		}

		solve(ar, arSum)
		i++
	}
}

func solve(ar []int, totalSum int) {
	leftSum := 0
	for i:=0; i<len(ar); i++ {
		rightSum := totalSum - leftSum - ar[i]
		if leftSum == rightSum {
			fmt.Print("YES\n")
			return
		}
		leftSum += ar[i]
	}

	fmt.Print("NO\n")
}
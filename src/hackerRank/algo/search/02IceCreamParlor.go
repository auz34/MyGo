/*
Sunny and Johnny together have M dollars they want to spend on ice cream. The parlor offers N flavors,
and they want to choose two flavors so that they end up spending the whole amount.

You are given the cost of these flavors. The cost of the iithth flavor is denoted by ccii. You have to display the indices of the two flavors whose sum is MM.

Input Format

The first line of the input contains TT; TT test cases follow.
Each test case follows the format detailed below: The first line contains MM. The second line contains NN. The third line contains NN space-separated integers denoting the price of each flavor. Here, the iithth integer denotes ccii.

Output Format

Output two integers, each of which is a valid index of a flavor. The lower index must be printed first. Indices are indexed from 11 to NN.

Constraints

1≤T≤501≤T≤50
2≤M≤100002≤M≤10000
2≤N≤100002≤N≤10000
1≤c1≤cii ≤10000,where i∈[1,N]≤10000,where i∈[1,N]
The prices of any two items may be the same and each test case has a unique solution.

Sample Input

2
4
5
1 4 5 3 2
4
4
2 2 4 3
Sample Output

1 4
1 2
Explanation

The sample input has two test cases.
For the 1st, the amount M = 4 and there are 5 flavors at the store. The flavors indexed at 1 and 4 sum up to 4.
For the 2nd test case, the amount M = 4 and the flavors indexed at 1 and 2 sum up to 4.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
)

type item struct {
	value int
	index int
}

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i:=0; i<testCases; {
		var money int
		fmt.Fscan(os.Stdin, &money)

		var n int
		fmt.Fscan(os.Stdin, &n)

		ar := make([]item, n)

		for j:=0; j<n; j++ {
			var v int
			fmt.Fscan(os.Stdin, &v)
			ar[j] = item{v, j + 1}
		}

		solve2(ar, money)
		i++
	}
}

func solve2(ar []item, money int) {
	sort(ar)
	for i:=0; i<len(ar); i++ {
		el1 := ar[i]
		if el1.value >= money {
			continue
		}

		ok, el2 := search(ar, money - el1.value, el1)
		if ok {
			if el1.index > el2.index {
				fmt.Printf("%v %v\n", el2.index, el1.index)
				return
			}

			fmt.Printf("%v %v\n", el1.index, el2.index)
			return
		}
	}
}

func sort(ar []item) {
	if len(ar) <=1 {
		return
	}

	rand.Seed(42)
	v := ar[rand.Intn(len(ar))]
	i, lt, gt := 0, 0, len(ar) - 1
	for i<=gt {
		comp := v.value - ar[i].value
		if comp > 0 {
			ar[i], ar[lt] = ar[lt], ar[i]
			lt++
			i++
		} else if comp < 0 {
			ar[i], ar[gt] = ar[gt], ar[i]
			gt--
		} else {
			i++
		}
	}

	sort(ar[0:lt])
	sort(ar[gt+1:])
}

func search(ar []item, value int, except item) (bool, item) {
	if len(ar) == 0 {
		return false, item{-1, -1}
	}

	if len(ar) == 1 {
		if ar[0].value == value && ar[0].index != except.index {
			return true, ar[0]
		}

		return false, ar[0]
	}

	mid := len(ar) / 2

	if ar[mid].value == value {
		if ar[mid].index != except.index {
			return true, ar[mid]
		}

		if mid + 1 < len(ar) && ar[mid+1].value == value {
			return true, ar[mid+1]
		}

		if mid>0 && ar[mid-1].value == value {
			return true, ar[mid-1]
		}
	}

	if ar[mid].value > value {
		return search(ar[0:mid], value, except)
	}

	if mid + 1 < len(ar) {
		return search(ar[mid+1:], value, except)
	}

	return false, ar[mid]
}
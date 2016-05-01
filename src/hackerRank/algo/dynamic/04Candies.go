/*Alice is a kindergarden teacher. She wants to give some candies to the children in her class.
All the children sit in a line ( their positions are fixed), and each  of them has a rating score according to his or
her performance in the class.  Alice wants to give at least 1 candy to each child. If two children sit next to each
other, then the one with the higher rating must get more candies. Alice wants to save money, so she needs to minimize
the total number of candies given to the children.

Input Format
The first line of the input is an integer N, the number of children in Alice's class. Each of the following N
lines contains an integer that indicates the rating of each child.
1 <= N <= 105
1 <= ratingi <= 105

Output Format
Output a single line containing the minimum number of candies Alice must buy.

Sample Input
3
1
2
2
Sample Output
4

Explanation
Here 1, 2, 2 is the rating. Note that when two children have equal rating, they are allowed to have different number
of candies. Hence optimal distribution will be 1, 2, 1.*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	rates := make([]int, n)
	for j:=0; j<n; j++ {
		var v int
		fmt.Fscan(os.Stdin, &v)
		rates[j] = v
	}

	fmt.Printf("%v", solve4(rates))
}

func solve4(rates []int) int {
	mem := make([]int, len(rates))
	mem[0]= 1

	for i:=1; i<len(rates); i++ {
		if rates[i] > rates[i-1] {
			mem[i] = mem[i-1] + 1
		} else {
			mem[i] = 1
		}
	}

	prevItem := 1
	for i:=len(rates) - 2; i>=0; i-- {
		next := 1
		if rates[i] > rates[i+1] {
			next = prevItem + 1
		}

		if mem[i]<next {
			mem[i] = next
		}

		prevItem = next
	}

	sum := 0
	for i:=0; i<len(mem); i++ {
		sum += mem[i]
	}

	return sum
}
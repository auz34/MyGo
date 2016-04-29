/*Emma is playing a new mobile game involving n clouds numbered from 0 to n−1. A player initially starts out on cloud c0,
and they must jump to cloud cn−1. In each step, she can jump from any cloud i to cloud i+1 or cloud i+2.

There are two types of clouds, ordinary clouds and thunderclouds. The game ends if Emma jumps onto a thundercloud,
but if she reaches the last cloud (i.e., cn−1), she wins the game!

Can you find the minimum number of jumps Emma must make to win the game? It is guaranteed that clouds c0 and cn−1 are
ordinary-clouds and it is always possible to win the game.

Input Format

The first line contains an integer, n (the total number of clouds).
The second line contains n space-separated binary integers describing clouds c0,c1,…,cn−1.

If ci=0, the ith cloud is an ordinary cloud.
If ci=1, the ith cloud is a thundercloud.

Constraints
2≤n≤100
ci∈{0,1}
c0=cn−1=0

Output Format

Print the minimum number of jumps needed to win the game.

Sample Input 0
7
0 0 1 0 0 1 0

Sample Output 0
4

Sample Input 1
6
0 0 0 0 1 0

Sample Output 1
3

Explanation

Sample Case 0:
Because c2 and c5 in our input are both 1, Emma must avoid c2 and c5. Bearing this in mind, she can win the game
with a minimum of 4 jumps


Sample Case 1:
The only thundercloud to avoid is c4. Emma can win the game in 3 jumps:*/

package main

import (
	"fmt"
	"os"
)

func main() {

	var n int
	fmt.Fscan(os.Stdin, &n)

	clouds := make([]int, n)
	for j:=0; j<n; j++ {
		var v int
		fmt.Fscan(os.Stdin, &v)
		clouds[j] = v
	}

	fmt.Printf("%v", solve1(clouds))

}

func min(a, b int) int {
	if a<b {
		return a
	}

	return b
}

func solve1(clouds []int) int {
	n := len(clouds)
	if n==1 {
		return 0
	}

	if n<=3 {
		return 1
	}


	mem := make([]int, n)
	mem[0] = 0
	mem[1] = 1


	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	for i:=2; i<n-1; i++ {
		if clouds[i] == 0 {
			mem[i] = min(mem[i-1], mem[i-2]) + 1
		} else {
			mem[i] = MaxInt
		}
	}


	return min(mem[len(mem) - 1], mem[len(mem)-2]) + 1
}


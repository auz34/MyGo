// A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3 steps at a time,
// implement a method to count how many possible ways the child can run up the stairs
package main

import "fmt"

func main() {
	fmt.Printf("Ways to hop 15 steps: %v \n", count(15))
}

func count(n int) int {
	mem := make([]int, n + 1)
	mem[0] = 0
	mem[1] = 1
	mem[2] = 2
	mem[3] = 4

	for i:=4; i<=n; i++ {
		mem[i] = mem[i-1] + mem[i-2] + mem[i-3]
	}

	return mem[n]
}
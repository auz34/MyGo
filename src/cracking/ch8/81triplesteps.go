// A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3 steps at a time,
// implement a method to count how many possible ways the child can run up the stairs
package main

import "fmt"

func main() {
	fmt.Printf("Ways to hop 15 steps: %v \n", count(15))
}

func count(n int) int {
	n1 := 1
	n2 := 2
	n3 := 4

	for i:=4; i<=n; i++ {
		n1, n2, n3 = n2, n3, n1+n2+n3
	}

	return n3
}
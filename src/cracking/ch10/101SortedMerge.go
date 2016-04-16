/* You are given two sorted arrays, A and B, and A has a large enough buffer at the end to hold B.
Write a method to merge B into A in sorted order.*/

package main

import "fmt"

func main() {
	a := []int {1, 2, 3, 4, 6, 8, 9, 10, 12, 0, 0, 0, 0, 0}
	b := []int {5, 7, 11, 15, 16}

	merge(a, b)
	for i:=0; i<len(a); i++ {
		fmt.Printf(" %v ", a[i])
	}
}

func merge(target, source []int) {
	sourcePoint := len(source) - 1
	targetPoint := len(target) - len(source) - 1
	assignPoint := len(target) - 1

	for sourcePoint>=0 {
		if target[targetPoint] > source[sourcePoint] {
			target[assignPoint] = target[targetPoint]
			targetPoint--
		} else {
			target[assignPoint] = source[sourcePoint]
			sourcePoint--
		}

		assignPoint--
	}
}

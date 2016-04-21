/* In an array of integers, a "peak" is an element which is greater than or equal to the adjacent integers and
a "valley" is an element which is less than or equal to the adjacent integers.
For example in the array {5, 8, 6, 2, 3, 4, 6} {8,6} are peaks amd {5,2} are valleys. Given an array of integers,
sort the array into an alternating sequence of peaks and valleys
EXAMPLE
Input: {5,3,1,2,3}
Output: {5,1,3,2,3}
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	input := []int {5,3,1,2,3,4,6,8,1}
	fmt.Print("Input: \n")
	printArray(input)
	sort(input)
	fmt.Print("Sorted Input: \n")
	printArray(input)
	fmt.Print("Answer: \n")
	alternatePeeks(input)
	printArray(input)
}

func printArray(ar []int) {
	for i:=0; i<len(ar); i++ {
		fmt.Printf(" %v ", ar[i])
	}

	fmt.Print("\n")
}

func sort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	rand.Seed(42)
	v := ar[rand.Intn(len(ar))]

	lt, i, gt := 0, 0, len(ar) - 1
	for i <= gt {
		comp := ar[i] - v
		if comp < 0 {
			ar[i], ar[lt] = ar[lt], ar[i]
			i++
			lt++
		} else if comp > 0 {
			ar[i], ar[gt] = ar[gt], ar[i]
			gt--
		} else {
			i++
		}
	}

	sort(ar[0:lt])
	sort(ar[gt+1:])
}

func alternatePeeks(ar []int) {
	mid := len(ar) / 2
	left := ar[0:mid]
	right := ar[mid:]

	for i:=0; i<len(left); i+=2 {
		left[i], right[i+1] = right[i+1], left[i]
	}
}



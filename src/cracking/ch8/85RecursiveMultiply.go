// Write a recursive function to multiply two positive integers without using * operator.
// You canuse addition, subtraction and bit shifting, but you should minimize the number of those operations.
package main

import "fmt"

func main() {
	a := uint(25)
	b := uint(17)
	fmt.Printf("%v*%v=%v\n", a, b, multiply(a,b))
}

func multiply(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b & uint(1) == 0 {
		a = a << 1
		b = b >> 1
	}

	sum := a
	b -= 1
	return sum + multiply(a, b)
}
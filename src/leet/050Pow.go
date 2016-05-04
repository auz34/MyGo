package main

import "fmt"

func main() {
	fmt.Print(myPow(float64(2), 3))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return float64(1)
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return float64(1) / myPow(x, -n)
	}

	temp := myPow(x, n / 2)
	if n % 2 == 0 {
		return temp*temp
	}

	return x*temp*temp
}
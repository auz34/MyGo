package main

import "fmt"

func main() {

	//s := "Hello world"

	fmt.Printf("%s\n", reverse("hello, world"))
	fmt.Printf("%s\n", reverse("he, world"))

	fmt.Printf("%s\n", reverse("01234"))
	fmt.Printf("%s\n", reverse("0123"))
}

func reverse(s string) string {
	var ars = []rune(s)

	fmt.Printf("Length of input %v\n", len(ars))

	var last = len(ars) - 1
	var mid =  len(ars) / 2
	fmt.Printf("len / 2 = %v\n", mid)

	for i := 0; i<mid; i++ {
		ars[i], ars[last - i] = ars[last - i], ars[i]
	}

	return string(ars)
}

// Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees. Can you do this in place?
package main

import "fmt"

func main() {
	array := [][]int {}

	array = append(array, []int{1, 2, 3, 4, 5})
	array = append(array, []int{6, 7, 8, 9, 10})
	array = append(array, []int{11, 12, 13, 14, 15})
	array = append(array, []int{16, 17, 18, 19, 20})
	array = append(array, []int{21, 22, 23, 24, 25})

	fmt.Print("Input array: \n")
	print(array)

	rotate(array)
	fmt.Print("Output array: \n")
	print(array)
}

func nextPlace(row, col, n int) (int, int) {
	return col, n-row-1
}

func rotate(ar [][]int) {
	n := len(ar)
	mid := n/2
	for i:=0; i<=mid; i++ {
		for j:=0; j<mid; j++ {
			r1, c1 := nextPlace(i, j, n)
			r2, c2 := nextPlace(r1, c1, n)
			r3, c3 := nextPlace(r2, c2, n)
			ar[i][j], ar[r1][c1], ar[r2][c2], ar[r3][c3] = ar[r3][c3], ar[i][j], ar[r1][c1], ar[r2][c2]
		}
	}
}

func print(ar [][]int) {
	for i:=0; i<len(ar); i++ {
		printRow(ar[i])
	}
}

func printRow(ar []int) {
	for i:=0; i<len(ar); i++ {
		fmt.Printf("%v\t", ar[i])
	}
	fmt.Print("\n")
}

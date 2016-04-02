// Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column is set to 0.
package main

import "fmt"

func main() {
	array := [][]int {}

	array = append(array, []int{1, 2, 3, 4, 5})
	array = append(array, []int{6, 0, 8, 0, 10})
	array = append(array, []int{11, 12, 13, 14, 15})
	array = append(array, []int{16, 17, 18, 19, 20})
	array = append(array, []int{21, 22, 23, 24, 0})

	fmt.Print("Input array: \n")
	print(array)

	setZeros(array)
	fmt.Print("Output array: \n")
	print(array)
}

func setZeros(ar [][]int) {
	cols := make(map[int]bool)
	rows := make(map[int]bool)
	for row:=0; row<len(ar); row++ {
		for col:=0; col<len(ar[row]); col++ {
			if ar[row][col] == 0 {
				cols[col] = true
				rows[row] = true
			}
		}
	}

	for k,_ := range rows {
		for col:=0; col<len(ar[k]); col++ {
			ar[k][col] = 0
		}
	}

	for k,_ := range cols {
		for row:=0; row<len(ar); row++ {
			ar[row][k] = 0
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

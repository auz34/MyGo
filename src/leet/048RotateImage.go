/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Follow up:
Could you do this in-place?
 */

package main

import "fmt"

func main() {
	matrix := [][]int {
		[]int {1, 2, 3, 4},
		[]int {5, 6, 7, 8},
		[]int {9, 10, 11, 12},
		[]int {13, 14, 15, 16},
	}

	/*matrix := [][]int {
		[]int {1, 2, 3},
		[]int {4, 5, 6 },
		[]int {7, 8, 9},
	}*/

	fmt.Print("Input:\n")
	printMatrix(matrix)
	rotate(matrix)
	fmt.Print("Output:\n")
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i:=0; i<len(matrix); i++ {
		fmt.Printf("%v\n", matrix[i])
	}
}

func getNewCoordinate(x, y, N int) (int, int) {
	return y, N - x - 1
}

func rotate(matrix [][]int)  {
	N := len(matrix)
	xmid, ymid := N / 2, N / 2
	if (N % 2 == 1) {
		xmid++
	}


	for i:=0; i<xmid; i++ {
		for j:=0; j<ymid; j++ {
			x1, y1 := i, j
			x2, y2 := getNewCoordinate(x1, y1, N)
			x3, y3 := getNewCoordinate(x2, y2, N)
			x4, y4 := getNewCoordinate(x3, y3, N)

			matrix[x1][y1], matrix[x2][y2], matrix[x3][y3], matrix[x4][y4] = matrix[x4][y4], matrix[x1][y1], matrix[x2][y2], matrix[x3][y3]
		}
	}
}

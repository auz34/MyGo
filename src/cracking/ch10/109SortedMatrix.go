// Given an MxN matrix in which each row and each column is sorted in ascending order,
// write a method to find an element.

package main

import "fmt"

func main() {
	matrix := make([][]int, 0)

	matrix = append(matrix, []int { 10, 20, 30, 40, 50 })
	matrix = append(matrix, []int { 11, 21, 31, 41, 51 })
	matrix = append(matrix, []int { 12, 22, 32, 42, 52 })
	matrix = append(matrix, []int { 13, 23, 33, 43, 53 })

	el13 := matrixSearch(matrix, 13, elIndex{0, 0}, elIndex{3, 4})
	fmt.Printf("Search result for 13 is {%v: %v} \n", el13.row, el13.col)

	el53 := matrixSearch(matrix, 53, elIndex{0, 0}, elIndex{3, 4})
	fmt.Printf("Search result for 53 is {%v: %v} \n", el53.row, el53.col)

	el40 := matrixSearch(matrix, 40, elIndex{0, 0}, elIndex{3, 4})
	fmt.Printf("Search result for 40 is {%v: %v} \n", el40.row, el40.col)

	el28 := matrixSearch(matrix, 28, elIndex{0, 0}, elIndex{3, 4})
	fmt.Printf("Search result for 28 is {%v: %v} \n", el28.row, el28.col)

}

type elIndex struct {
	row int
	col int
}

func rowSearch(matrix [][]int, num, row, left, right int) elIndex {
	if left > right {
		return elIndex{-1, -1}
	}

	mid := (left + right) / 2
	if matrix[row][mid] == num {
		return elIndex{row, mid}
	}

	if matrix[row][mid] > num {
		return rowSearch(matrix, num, row, left, mid - 1)
	}

	return rowSearch(matrix, num, row, mid+1, right)
}

func colSearch(matrix [][]int, num, col, top, bottom int) elIndex {
	if top > bottom {
		return elIndex{-1, -1}
	}

	mid := (top + bottom) / 2

	if matrix[mid][col] == num {
		return elIndex{mid, col}
	}

	if matrix[mid][col] > num {
		return colSearch(matrix, num, col, top, mid - 1)
	}

	return colSearch(matrix, num, col, mid + 1, bottom)
}

func matrixSearch(matrix [][]int, num int, tl, br elIndex) elIndex {
	if br.row - tl.row <= 1 {
		for row := tl.row; row <= br.row; row++ {
			result := rowSearch(matrix, num, row, tl.col, br.col)
			if result.col != -1 {
				return result
			}
		}

		return elIndex{-1, -1}
	}

	if br.col - tl.col <= 1 {
		for col := tl.col; col <= br.col; col++ {
			result := colSearch(matrix, num, col, tl.row, br.row)
			if result.col != -1 {
				return result
			}
		}

		return elIndex{-1, -1}
	}

	mid := elIndex{ (tl.row + br.row) / 2, (tl.col + br.col) / 2 }

	midNum := matrix[mid.row][mid.col]
	if  midNum == num {
		return mid
	}

	if midNum < num {
		// search right (not including mid column)
		result := matrixSearch(matrix, num, elIndex{tl.row, mid.col+1}, br)
		if result.row != -1 {
			return result
		}

		// search left and bottom (including mid column)
		return matrixSearch(matrix, num, elIndex{ mid.row + 1, tl.col}, elIndex{br.row, mid.col})
	}

	// search left
	result := matrixSearch(matrix, num, tl, elIndex{ br.row, mid.col })
	if result.row != -1 {
		return result
	}

	// search right and top
	return matrixSearch(matrix, num, elIndex{tl.row, mid.col}, elIndex{mid.row - 1, br.col})
}

/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and
is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all
surrounded by water.

Example 1:

11110
11010
11000
00000
Answer: 1

Example 2:

11000
11000
00100
00011
Answer: 3
 */
package main

import "fmt"

func main() {
	input := make([][]byte, 0)
	input = append(input, []byte{1, 1, 1, 1, 0})
	input = append(input, []byte{1, 1, 0, 1, 0})
	input = append(input, []byte{1, 1, 0, 0, 0})
	input = append(input, []byte{0, 0, 0, 0, 0})

	fmt.Printf("Input #1: %v \n", numIslands(input))

	input2 := make([][]byte, 0)
	input2 = append(input2, []byte{0, 1, 1, 1, 0})
	input2 = append(input2, []byte{1, 1, 0, 1, 0})
	input2 = append(input2, []byte{1, 1, 0, 0, 0})
	input2 = append(input2, []byte{0, 0, 0, 0, 0})

	fmt.Printf("Input #2: %v \n", numIslands(input2))

	input3 := make([][]byte, 0)
	input3 = append(input3, []byte{1})


	fmt.Printf("Input #3: %v \n", numIslands(input3))
}

func numIslands(grid [][]byte) int {
	id := 0
	mem := make([][]int, len(grid))
	joins := make(map[int]int)
	for i:=0; i<len(grid); i++ {
		row := grid[i]
		mem[i] = make([]int, len(row))
		for j:=0; j<len(row); j++ {
			if grid[i][j] == 1 {
				up := 0
				if i != 0 {
					up = mem[i-1][j]
				}

				left := 0
				if j != 0 {
					left = mem[i][j-1]
				}



				if up == 0 && left == 0 {
					id++
					mem[i][j] = id
					continue
				}

				if up == 0 && left != 0 {
					mem[i][j] = left
					continue
				}

				if up != 0 && left == 0 {
					mem[i][j] = up
					continue
				}

				if up != left {
					joins[left] = up
				}

				mem[i][j] = up
			}
		}

	}

	return id - len(joins)
}
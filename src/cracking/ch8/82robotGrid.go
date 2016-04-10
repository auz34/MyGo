// Imagine a robot sitting on the upper left corner of grid with r rows and
// c columns. The robot can only move in two directions, right and down, but certain cells are
// "off limits" such that robot cannot step on them. Design an algorithm to find a path for the
// robot from the top left to the bottom right
package main

import "fmt"

func main() {
	placeMap := make([][]bool, 0)
	t := true
	f := false
	placeMap = append(placeMap, []bool { t, t, f, t, t})
	placeMap = append(placeMap, []bool { t, t, f, t, t})
	placeMap = append(placeMap, []bool { f, t, t, t, t})
	placeMap = append(placeMap, []bool { f, t, t, t, t})
	findPath(placeMap, 4, 5)
}

func findPath(placeMap [][]bool, r, c int) {
	steps := make([][]bool, r)
	steps[0] = make([]bool, c)
	steps[0][0]= placeMap[0][0]
	for i:=0; i<r; i++ {
		if steps[i] == nil {
			steps[i] = make([]bool, c)
		}

		for j:=0; j<c; j++ {
			if placeMap[i][j] && ((i>0 && steps[i-1][j]) || (j>0 && steps[i][j-1])) {
				steps[i][j] = true
			}
		}
	}

	if !steps[r-1][c-1] {
		fmt.Print("top bottom point is unreachable")
		return
	}

	path := fmt.Sprintf("{%v, %v}", r-1, c-1)
	for i, j := r-1, c-1; i>0 || j>0; {
		if i>0 && steps[i-1][j] {
			i--
			path = fmt.Sprintf("{%v, %v} -> %v", i, j, path)
		} else {
			j--
			path = fmt.Sprintf("{%v, %v} -> %v", i, j, path)
		}
	}

	fmt.Printf("Solution: %v \n", path)
}

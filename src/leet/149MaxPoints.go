/*
Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
*/

package main

import (
	"fmt"
	//"math"
)

type Point struct {
	X int
	Y int
}

func main() {
	ar := []Point{
		Point{1,1},
		Point{7,2},
		Point{2,2},
		Point{4,3},
		Point{3,3},
	}

	fmt.Printf("%v\n", maxPoints(ar))
}

func max149(a,b int) int {
	if a > b {
		return a
	}

	return b
}

func isEqual(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func isVertical(p1, p2 Point) bool {
	return p1.X == p2.X
}

func slope(p1, p2 Point) float64 {
	real := float64(p1.Y - p2.Y)/float64(p1.X - p2.X)
	return real
	//return int(math.Ceil(real * float64(1000000)))
}

func maxPoints(points []Point) int {
	if len(points) == 0 {
		return 0
	}

	if len(points) == 1 {
		return 1
	}

	if len(points) == 2 {
		return 2
	}

	dict := make(map[float64]int)
	result := 1
	duplicates := 0
	p1 := points[0]
	for i:=1; i<len(points); i++ {
		if isEqual(p1, points[i]) {
			duplicates++
			result++
			continue
		}

		if isVertical(p1, points[i]) {
			result++
		} else {
			sl := slope(p1, points[i])
			dict[sl]++
		}
	}

	for _, v := range dict {
		result = max149(result, v + 1 + duplicates)
	}

	return max149(result, maxPoints(points[1:]))
}

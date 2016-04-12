// Write an all algorithm to print all ways of arranging eight queens on an 8x8 chess board so that none of
// them share the same rom column or diagonal. In this case, "diagonal" means all diagonals, not just the two
// that bisect the board

package main

import "fmt"

type position struct {
	row int
	col int
}

func abs(x int) int {
	if x>0 {
		return x
	}

	return -x
}

func (self *position) beatsPos(pos position) bool {
	if self.row == pos.row || self.col == pos.col {
		return true
	}

	return abs(self.row-pos.row) == abs(self.col-pos.col)
}


func main() {
	fmt.Print("Search for solutions:\n")
	solution := make([]position, 0)
	search(solution)
}

func search(solution []position) {
	l := len(solution)
	if l==8 {
		printSolution(solution)
		return
	}

	for i:=0; i<8; i++ {
		candidate := position{l, i}
		validCandidate := true
		for j:=0; validCandidate && j<len(solution); j++ {
			validCandidate = validCandidate && !candidate.beatsPos(solution[j])
		}

		if validCandidate {
			solution = append(solution, candidate)
			search(solution)
			solution = solution[:l]
		}
	}
}

func printSolution(solution []position) {
	for _,v := range solution {
		fmt.Printf(" {%v, %v} ", v.row, v.col);
	}

	fmt.Print("\n")
}
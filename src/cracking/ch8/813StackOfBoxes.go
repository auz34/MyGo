// You have a stack of n boxes with width wi, heights hi, and depths di. The boxes cannot be rotated and can only
// be stacked on top of the another if each box in the stack is strictly larger than box above it in width, height,
// and depth. Implement a method to compute the height of the tallest possible stack.
// The height of stack is the sum of the heights of each box.

package main

import (
	"fmt"
	"sort"
)

type box struct {
	width int
	height int
	depth int
}

type by func (b1, b2 *box) bool

func (by by) sort(boxes []box) {
	ps := &boxSorter{
		boxes: boxes,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type boxSorter struct {
	boxes []box
	by by
}

// Len is part of sort.Interface.
func (s *boxSorter) Len() int {
	return len(s.boxes)
}

// Swap is part of sort.Interface.
func (s *boxSorter) Swap(i, j int) {
	s.boxes[i], s.boxes[j] = s.boxes[j], s.boxes[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *boxSorter) Less(i, j int) bool {
	return s.by(&s.boxes[i], &s.boxes[j])
}

func main() {
	boxes := make([]box, 0)
	boxes = append(boxes, box{1, 7, 1})
	boxes = append(boxes, box{6, 6, 6})
	boxes = append(boxes, box{1, 1, 1})
	boxes = append(boxes, box{2, 2, 2})
	boxes = append(boxes, box{3, 3, 3})
	boxes = append(boxes, box{4, 4, 4})
	boxes = append(boxes, box{5, 5, 5})

	findTallestHeight(boxes)

	fmt.Printf("tallest possible height is %v", findTallestHeight(boxes))
}

func findTallestHeight(boxes []box) int {
	height := func (b1, b2 *box) bool {
		return b1.height > b2.height
	}

	all := func (b1, b2 *box) bool {
		return b1.height > b2.height &&
			   b1.width > b2.width &&
		       b1.depth > b2.depth
	}

	by(height).sort(boxes)

	heights := make([]int, len(boxes))
	result := 0
	for i:=0; i<len(boxes); i++ {
		maxPrev := 0
		for j:=0; j<i; j++ {
			if all(&boxes[j], &boxes[i]) && maxPrev < heights[j] {
				maxPrev = heights[j]
			}
		}

		heights[i] = maxPrev + boxes[i].height
		if result < heights[i] {
			result = heights[i]
		}
	}

	return result

}

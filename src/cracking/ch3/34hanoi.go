/*In the classic problem of the Towers of Hanoi, you have 3 rods and N disks of different
sizes which can slide onto any tower. The puzzle starts with disks sorted in ascending
order of size from top to bottom (e.g., each disk sits on top of an even larger one). You
have the following constraints:
(A) Only one disk can be moved at a time.
(B) A disk is slid off the top of one rod onto the next rod.
(C) A disk can only be placed on top of a larger disk.
Write a program to move the disks from the first rod to the last using Stacks*/

package main

import (
	"fmt"
	"cracking/ch3/basics"
)

type hanoiPuzzle struct {
	numberOfDisks int
	first *basics.Stack
	second *basics.Stack
	third *basics.Stack
}

func NewHanoiPuzzle(numberOfDisks int) *hanoiPuzzle {
	first, second, third := basics.NewStack(), basics.NewStack(), basics.NewStack()

	for i:=numberOfDisks; i>0; i-- {
		first.Push(i)
	}

	return &hanoiPuzzle{ numberOfDisks, first, second, third }
}

func (self *hanoiPuzzle) getStackByNum(num int) *basics.Stack {
	switch num {
		case 1: return self.first
		case 2: return self.second
		case 3: return self.third
	}

	return nil
}

func (self *hanoiPuzzle) getMissingStack(first, second int) int {
	if first == 1 || second == 1 {
		if first == 2 || second == 2 {
			return 3
		}

		return 2
	}

	return 1
}

func (self *hanoiPuzzle) move(source, target int) {
	sourceStack := self.getStackByNum(source)
	targetStack := self.getStackByNum(target)

	d, _ := sourceStack.Pop()
	targetStack.Push(d)
	fmt.Printf("D%v #%v -> #%v \n", d, source, target)
}

func (self *hanoiPuzzle) recursiveSolve(numberOfDisks, source, target int) {
	if numberOfDisks == 1 {
		self.move(source, target)
		return
	}

	proxy := self.getMissingStack(source, target)
	self.recursiveSolve(numberOfDisks - 1, source, proxy)
	self.move(source, target)
	self.recursiveSolve(numberOfDisks - 1, proxy, target)
}

func (self *hanoiPuzzle) solve() {
	self.recursiveSolve(self.numberOfDisks, 1, 3)
}

func main() {
	puzzle := NewHanoiPuzzle(3)
	puzzle.solve()
}
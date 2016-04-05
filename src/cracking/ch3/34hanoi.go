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
	"errors"
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

func (self *hanoiPuzzle) move(source, target int) {
	sourceStack := self.getStackByNum(source)
	targetStack := self.getStackByNum(target)

	d, _ := sourceStack.Pop()
	targetStack.Push(d)
	fmt.Printf("Moved disk #%v from tower #%v to tower #%v", d, source, target)
}

func (self *hanoiPuzzle) reverseSolve(numberOfDisks int) {
	if numberOfDisks == 1 {
		self.move(3, 1)
	}

	self.reverseSolve(numberOfDisks - 1)
	self.move(3, 2)
}

func (self *hanoiPuzzle) recursiveSolve(numberOfDisks int) {
	if numberOfDisks == 1 {
		self.move(1, 3)
		return
	}

	self.recursiveSolve(numberOfDisks - 1)
	self.move(1, 2)

}

func (self *hanoiPuzzle) solve() {
	self.recursiveSolve(self.numberOfDisks)
}

func main() {
	puzzle := NewHanoiPuzzle(5)
	puzzle.solve()
}
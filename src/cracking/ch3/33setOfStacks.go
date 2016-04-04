/* Imagine a (literal) stack of plates If the stack gets too high, it might topple
Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold
Implement a data structure SetOfStacks that mimics this SetOf- Stacks should be composed of several stacks,
and should create a new stack once the previous one exceeds capacity SetOfStacks push() and SetOfStacks pop()
should behave identically to a single stack (that is, pop() should return the same values as it would if
there were just a single stack)
FOLLOW UP
Implement a function popAt(int index) which performs a pop operation on a speci c sub-stack*/

package main

import (
	"fmt"
	"cracking/ch3/basics"
	"errors"
)

type setOfStacks struct {
	maxHeight int
	set []*basics.Stack
}

func NewSetOfStacks(maxHeight int) *setOfStacks {
	return &setOfStacks{ maxHeight, make([]*basics.Stack, 0) }
}

func (self *setOfStacks) Push (v int) {
	l := len(self.set)
	if l == 0 {
		self.set = append(self.set, basics.NewStack())
	}

	l = len(self.set)
	lastStack := self.set[l-1]

	if lastStack.Count() == self.maxHeight {
		lastStack = basics.NewStack()
		self.set = append(self.set, lastStack)
	}

	lastStack.Push(v)
}

func (self *setOfStacks) Pop() (int, error) {
	l := len(self.set)
	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	lastStack := self.set[l-1]
	res, _ := lastStack.Pop()

	if lastStack.Count() == 0 {
		self.set = self.set[:l-1]
	}

	return res, nil
}

func main() {
	stack := NewSetOfStacks(4)
	stack.Push(1);
	stack.Push(2);
	stack.Push(3); stack.Push(4)
	stack.Push(5); stack.Push(6); stack.Push(7); stack.Push(8)
	stack.Push(9); stack.Push(10); stack.Push(11); stack.Push(12)
	stack.Push(13); stack.Push(14); stack.Push(15)

	v, e := stack.Pop()
	for e == nil {
		fmt.Printf(" %v ->", v)
		v, e = stack.Pop()
	}

}
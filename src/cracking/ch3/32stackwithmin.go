// How would you design a stack which, in addition to push and pop, also has a function min which
// returns the minimum element? Push, pop and min should all operate in O(1) time
package main

import (
	"fmt"
	"cracking/ch3/basics"
)

type stackWithMin struct {
	stack *basics.Stack
	minStack *basics.Stack
}

func NewStackWithMin() *stackWithMin {
	return &stackWithMin{ basics.NewStack(), basics.NewStack() }
}

func (self *stackWithMin) Push(v int) {
	if self.stack.Count() == 0 {
		self.minStack.Push(v)
	} else {
		oldMin, _ := self.minStack.Peek()
		if oldMin < v {
			self.minStack.Push(oldMin)
		} else {
			self.minStack.Push(v)
		}
	}

	self.stack.Push(v)
}

func (self *stackWithMin) Pop() (int, error) {
	self.minStack.Pop()
	return self.stack.Pop()
}

func (self *stackWithMin) Min() (int, error) {
	return self.minStack.Peek()
}

func main() {
	stack := NewStackWithMin()
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(1)
	min, _ := stack.Min()
	fmt.Printf("%v should be equal to 1 at this stage \n", min)
	stack.Pop()
	min, _ = stack.Min()
	fmt.Printf("%v should be equal to 4 at this stage \n", min)
}

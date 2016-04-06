package main

import (
	"fmt"
	"cracking/ch3/basics"
)

type myQueue struct {
	first *basics.Stack
	second *basics.Stack
}

func NewQueue() *myQueue {
	return &myQueue{ basics.NewStack(), basics.NewStack() }
}

func (self *myQueue) push(v int) {
	self.first.Push(v)
}

func (self *myQueue) pop() (int, error) {
	if self.second.Count() != 0 {
		return self.second.Pop()
	}

	for self.first.Count() != 0 {
		v, e := self.first.Pop()
		if e != nil {
			return 0, e
		}

		self.second.Push(v)
	}

	return self.second.Pop()
}

func main() {
	q := NewQueue()
	q.push(1)
	q.push(2)
	q.push(3)
	v, _ := q.pop()
	fmt.Printf("%v \n", v)
	q.push(4)

	v, _ = q.pop()
	fmt.Printf("%v \n", v)

	v, _ = q.pop()
	fmt.Printf("%v \n", v)

	v, _ = q.pop()
	fmt.Printf("%v \n", v)
}

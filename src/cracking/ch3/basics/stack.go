package basics

import "errors"

type stack struct {
	internal []int
}

func NewStack() *stack {
	return &stack{ make([]int, 0)}
}

func (self *stack) Push(v int) {
	self.internal = append(self.internal, v)
}

func (self *stack) Pop() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	self.internal = self.internal[:l-1]

	return res, nil
}

func (self *stack) Peek() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	return res, nil
}



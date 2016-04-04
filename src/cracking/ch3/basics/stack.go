package basics

import "errors"

type Stack struct {
	internal []int
}

func NewStack() *Stack {
	return &Stack{ make([]int, 0)}
}

func (self *Stack) Count() int {
	return len(self.internal)
}

func (self *Stack) Push(v int) {
	self.internal = append(self.internal, v)
}

func (self *Stack) Pop() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	self.internal = self.internal[:l-1]

	return res, nil
}

func (self *Stack) Peek() (int, error) {
	l := len(self.internal)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := self.internal[l-1]
	return res, nil
}



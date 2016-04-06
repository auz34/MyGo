/*Write a program to sort a stack in ascending order. You should not make any assumptions
about how the stack is implemented. The following are the only functions that
should be used to write this program: push | pop | peek | isEmpty.*/

package main

import (
	"fmt"
	"cracking/ch3/basics"
)

func sort(st *basics.Stack) {
	aux := basics.NewStack()

	for st.Count() != 0 {
		v, _ := st.Pop()
		if aux.Count() == 0 {
			aux.Push(v)
			continue
		}

		p, _ := aux.Peek()
		for v < p {
			v1, _ := aux.Pop()
			st.Push(v1)
			if aux.Count() == 0 {
				p = 0
			} else {
				p, _ = aux.Peek()
			}
		}

		aux.Push(v)
	}

	for aux.Count() != 0 {
		v, _ := aux.Pop()
		st.Push(v)
	}
}

func main() {
	stack := basics.NewStack()
	stack.Push(5); stack.Push(2); stack.Push(4); stack.Push(8);

	sort(stack)
	for stack.Count() != 0 {
		v, _ := stack.Pop()
		fmt.Printf(" %v ", v)
	}

	fmt.Print("\n")
}

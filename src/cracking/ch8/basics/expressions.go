package basics

import (
	//"fmt"
	"unicode/utf8"
)

func InfixToPostfix(expr string) string {
	s := NewStack()
	result := ""
	for i:=0; i<len(expr); {
		r, w := utf8.DecodeRuneInString(expr[i:])
		i += w

		switch r {
		case '0', '1':
			result += string(r)
		case '(':
			s.Push(r)
		case ')':
			for p, _ := s.Peek(); p != '('; {
				r, _ := s.Pop()
				result += string(r)
				p, _ = s.Peek()
			}

			s.Pop() // remove "("
		case '&', '|', '^':
			// if we had operators with different precedence we would have more complex code here
			s.Push(r)
		}
	}

	for s.Count() != 0 {
		r, _ := s.Pop()
		result += string(r)
	}

	return result
}

func EvalPostfix(expr string) bool {
	return false
}



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

func applyOp(r1, r2, op rune) rune {
	b1 := r1 == '1'
	b2 := r2 == '1'
	res := false
	switch op {
	case '&':
		res = b1 && b2
	case '|':
		res = b1 || b2
	case '^':
		res = (b1 || b2) && !(b1 && b2)
	}

	if res {
		return '1'
	}

	return '0'
}

func EvalPostfix(expr string) bool {
	s := NewStack()
	for i:=0; i<len(expr); {
		r, w := utf8.DecodeRuneInString(expr[i:])
		i += w
		switch r {
		case '0', '1': s.Push(r)
		case '&', '|', '^':
			r1,_ := s.Pop()
			r2,_ := s.Pop()
			r3 := applyOp(r1, r2, r)
			s.Push(r3)
		}
	}

	res, _ := s.Pop()
	if res == '1' {
		return true
	}

	return false
}
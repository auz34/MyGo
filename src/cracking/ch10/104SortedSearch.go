/* You are given an array-like data structure Listy which lacks a size method. It does, however, have an elementAt(i)
method that returns the element at index i in O(1) time. If i beyond the bounds of the data structure , it returns -1.
(For this reason, the data structure only supports positive integers.) Given a Listy which contains sorted positive
integers, find the index at which an element x occurs. If x occurs multiple times, you may return any index
*/

package main

import "fmt"

type listy struct {
	internal []int
}

func (self *listy) elementAt(index int) int {
	if index < 0 || index > len(self.internal) - 1 {
		return -1
	}

	return self.internal[index]
}

func createListy(ar []int) *listy {
	return &listy{ ar }
}

func main() {
	input := createListy([]int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 })


	fmt.Printf("Index of 10 is %v \n", listySearch(input, 10))
	fmt.Printf("Index of 7 is %v \n", listySearch(input, 7))
	fmt.Printf("Index of 2 is %v \n", listySearch(input, 2))
}

func binaryListySearch(list *listy, num, left, right int) int {
	if left>right {
		return -1
	}

	mid := (left + right) / 2
	if list.elementAt(mid) == num {
		return mid
	} else if list.elementAt(mid) < num {
		return binaryListySearch(list, num, left, mid - 1)
	} else {
		return binaryListySearch(list, num, mid + 1, right)
	}
}

func findRange(list *listy, num, left, right int) (int, int) {
	l := list.elementAt(left)
	r := list.elementAt(right)

	if l==-1 || l>num {
		return 10, 8
	}

	if l == num {
		return left, left
	}

	if r == num {
		return right, right
	}

	if r == -1 {
		return findRange(list, num, left, (left+right)/2)
	}

	if r > num {
		return left, right
	}

	return findRange(list, num, right, right * 2)
}


func listySearch(list *listy, num int) int {
	left, right := findRange(list, num, 0, 1)
	fmt.Printf("Found range %v:%v \n", left, right)
	return binaryListySearch(list, num, left, right)
}
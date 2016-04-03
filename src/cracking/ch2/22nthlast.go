// Implement an algorithm to  nd the nth to last element of a singly linked list
package main

import "fmt"

func main() {
	ar := []int { 1, 8, 4, 9, 8, 12, 16 }
	root := generateList(ar)
	fmt.Printf("Input List: \n")
	printList(root)
	fmt.Printf("Second from the end = %v \n", root.GetNthFromEnd(1).Value)
}

type Node struct {
	Value int
	Next *Node
}

func (self *Node) Count() int {
	curNode := self
	cnt := 0
	for curNode != nil {
		cnt++
		curNode = curNode.Next
	}

	return cnt
}

func (self *Node) GetNth(n int) *Node {
	curNode := self
	cnt := 0

	for curNode != nil {
		cnt++
		if (cnt == n) {
			return curNode
		}
		curNode = curNode.Next
	}

	return nil
}

func (self *Node) GetNthFromEnd(n int) *Node {
	cnt := self.Count()
	return self.GetNth(cnt - n)
}

func generateList(ar []int) *Node {
	var prev, root *Node = nil, nil
	for _, v := range ar {
		node := Node { v, nil }
		if prev != nil {
			prev.Next = &node
		} else {
			root = &node
		}

		prev = &node
	}

	return root
}

func printList(root *Node) {
	var curNode *Node = root

	for curNode != nil {
		fmt.Printf("%v -> ", curNode.Value)
		curNode = curNode.Next
	}

	fmt.Print("\n")
}
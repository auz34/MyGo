// Implement an algorithm to  nd the nth to last element of a singly linked list
package main

import "fmt"

func main() {
	ar := []int { 1, 8, 4, 9, 8, 12, 16 }
	root := generateList(ar)
	fmt.Printf("Input List: \n")
	printList(root)
	res, _ := root.GetNthFromEnd(1)
	fmt.Printf("Second from the end = %v \n", res.Value)
}

type Node struct {
	Value int
	Next *Node
}

func (self *Node) GetNthFromEnd(n int) (*Node, int) {
	ndx := 0
	var res *Node = nil
	if self.Next != nil {
		res, ndx = self.Next.GetNthFromEnd(n)
		ndx++
	}

	if (ndx == n) {
		res = self
	}

	return res, ndx
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
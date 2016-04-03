// Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func (self *Node) removeNextFromList() {
	self.Next = self.Next.Next
}

func (self *Node) removeSubsequentDuplicates() {
	curNode := self
	for curNode != nil {
		if curNode.Next != nil && curNode.Next.Value == self.Value {
			curNode.removeNextFromList()
		}

		curNode = curNode.Next
	}
}

func main() {
	ar := []int { 1, 8, 4, 9, 8, 12, 16 }
	root := generateList(ar)
	fmt.Printf("Input List: \n")
	printList(root)
	removeDuplicates(root)
	fmt.Printf("Output List: \n")
	printList(root)

}

func removeDuplicates(root *Node) {
	curNode := root
	for curNode != nil {
		curNode.removeSubsequentDuplicates()
		curNode = curNode.Next
	}
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
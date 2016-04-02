// Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func main() {
	ar := []int { 1, 8, 4, 9, 8, 12, 16 }
	root := generateList(ar)
	printList(root)
}

func generateList(ar []int) *Node {
	var prev, root *Node = nil, nil
	for v, _ := range ar {
		node := Node { v, nil }
		if prev != nil {
			node.Next = &node
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
		fmt.Printf("%v -> ", *curNode.Value)
		curNode = *curNode.Next
	}
}
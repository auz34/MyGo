// Implement an algorithm to delete a node in the middle of a single linked list, given only access to that node
// EXAMPLE
// Input: the node ‘c’ from the linked list a->b->c->d->e
// Result: nothing is returned, but the new linked list looks like a->b->d->e
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
	fmt.Printf("Input List: \n")
	printList(root)
	removeNode(root.Next.Next.Next)

	fmt.Printf("Output List: \n")
	printList(root)

}

func removeNode(node *Node) {
	if (node.Next == nil) {
		panic("can not complete the task")
	}

	node.Value = node.Next.Value
	if (node.Next.Next == nil) {
		node.Next = nil
	} else {
		removeNode(node.Next)
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

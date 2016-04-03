// You have two numbers represented by a linked list, where each node contains a sin- gle digit
// The digits are stored in reverse order, such that the 1’s digit is at the head of the list
// Write a function that adds the two numbers and returns the sum as a linked list
// EXAMPLE
// Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
// Output: 8 -> 0 -> 8
package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func main() {
	n1 := []int { 3, 1, 5 }
	n2 := []int { 5, 9, 2 }
	root1 := generateList(n1)
	root2 := generateList(n2)
	fmt.Printf("First number: \n")
	printList(root1)
	fmt.Printf("Second number: \n")
	printList(root2)
	fmt.Printf("Sum of two: \n")
	printList(getSum(root1, root2))
}

func getSum(n1, n2 *Node) *Node {
	curNode1 := n1
	curNode2 := n2

	var resNode, lastNode *Node = nil, nil

	mem := 0
	for curNode1 != nil && curNode2 != nil {
		v1 := 0
		if curNode1 != nil {
			v1 = curNode1.Value
		}

		v2 := 0
		if curNode1 != nil {
			v2 = curNode2.Value
		}

		value := (v1 + v2 + mem) % 10
		mem = (v1 + v2 + mem) / 10

		node := Node { value, nil }
		if (resNode == nil) {
			resNode = &node
		}

		if (lastNode != nil) {
			lastNode.Next = &node
		}

		if curNode1 != nil {
			curNode1 = curNode1.Next
		}

		if curNode2 != nil {
			curNode2 = curNode2.Next
		}

		lastNode = &node
	}

	return resNode
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

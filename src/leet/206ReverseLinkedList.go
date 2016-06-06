package main

import "fmt"

func main() {
	list := generateList206([]int {1, 3, 5, 7, 10, 15, 16, 17})
	fmt.Print("Input:\n")
	printList206(list)
	fmt.Print("Output: \n")
	result := reverseList(list)
	printList206(result)
}

func generateList206(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}

	root := &ListNode{items[0], nil}
	prev := root

	for i:=1; i<len(items); i++ {
		next := &ListNode{items[i], nil}
		prev.Next = next
		prev = next
	}

	return root
}

func printList206(node *ListNode) {
	for node != nil {
		fmt.Printf(" %v ", node.Val)
		node = node.Next
	}

	fmt.Print("\n")
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	nextNode := head
	var prevNode *ListNode = nil

	for nextNode != nil {
		curNode := nextNode
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
	}

	return prevNode
}

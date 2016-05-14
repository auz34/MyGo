/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the
nodes of the first two lists.
 */
package main

import "fmt"

func main() {
	list1 := generateList21([]int {1, 4, 5, 8, 11})
	list2 := generateList21([]int {1, 3, 5, 7, 10, 15, 16, 17})
	fmt.Print("Input:\n")
	printList21(list1)
	printList21(list2)
	fmt.Print("Output: \n")
	result := mergeTwoLists(list1, list2)
	printList21(result)
}

func generateList21(items []int) *ListNode {
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

func printList21(node *ListNode) {
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev, root *ListNode
	for l1 != nil || l2 != nil {
		var next int
		if l1 == nil || (l2 != nil && l2.Val < l1.Val) {
			next = l2.Val
			l2 = l2.Next
		} else {
			next = l1.Val
			l1 = l1.Next
		}

		nextNode := &ListNode{next, nil}
		if prev !=nil {
			prev.Next = nextNode
		} else {
			root = nextNode
		}

		prev = nextNode
	}

	return root
}
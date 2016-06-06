/*
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.
 */

package main

import "fmt"

func main() {
	list := generateList92([]int {1, 3, 5, 7, 10, 15, 16, 17})
	fmt.Print("Input:\n")
	printList92(list)
	fmt.Print("Output: \n")
	result := reverseBetween(list, 5, 8)
	printList92(result)
}

func generateList92(items []int) *ListNode {
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

func printList92(node *ListNode) {
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	nextNode := head
	var prevNode *ListNode = nil
	i := 1

	for i < m {
		prevNode = nextNode
		nextNode = nextNode.Next
		i++
	}

	preFirst := prevNode
	first := nextNode

	prevNode = nil
	for i:=m; i<=n; i++ {
		curNode := nextNode
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
	}

	first.Next = nextNode
	if preFirst != nil {
		preFirst.Next = prevNode
		return head
	}

	return prevNode
}
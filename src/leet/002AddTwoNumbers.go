/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

package main

import "fmt"

func main() {
	list1 := generateList002([]int {2, 4, 3})
	list2 := generateList002([]int {5, 6, 4})
	fmt.Print("Input:\n")
	printList002(list1)
	printList002(list2)
	fmt.Print("Output: \n")
	result := addTwoNumbers(list1, list2)
	printList002(result)
}

func generateList002(items []int) *ListNode {
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

func printList002(node *ListNode) {
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

func getVal002(node *ListNode) int {
	if node != nil {
		return node.Val
	}

	return 0
}

func getNext(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if (l1 == nil) {
		return l2
	}

	if (l2 == nil) {
		return l1
	}



	carry := (l1.Val + l2.Val) / 10
	answer := &ListNode{(l1.Val + l2.Val) % 10, nil}
	last := answer

	p1, p2 := l1.Next, l2.Next

	for p1 != nil || p2 != nil {
		d1 := getVal002(p1)
		d2 := getVal002(p2)

		cur := &ListNode{(d1 + d2 + carry) % 10, nil}
		carry = (d1 + d2 + carry) / 10

		last.Next = cur
		last = cur

		p1 = getNext(p1)
		p2 = getNext(p2)
	}

	if carry > 0 {
		last.Next = &ListNode{carry, nil}
	}

	return answer
}

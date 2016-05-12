/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
 */

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list1 := generateList([]int {1, 4, 5, 8, 11})
	list2 := generateList([]int {1, 3, 5, 7, 10})
	list3 := generateList([]int {2, 9, 10, 15})
	fmt.Print("Input:\n")
	printList(list1)
	printList(list2)
	printList(list3)
	fmt.Print("Output: \n")
	result := mergeKLists([]*ListNode {list1, list2, list3})
	printList(result)

}

func generateList(items []int) *ListNode {
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

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf(" %v ", node.Val)
		node = node.Next
	}

	fmt.Print("\n")
}

const MaxUint = ^uint(0)
//const MinUint = 0
const MaxInt = int(MaxUint >> 1)
//const MinInt = -MaxInt - 1

func nextNode(lists []*ListNode) *ListNode {
	minValue := MaxInt
	minNdx := -1

	for i:=0; i<len(lists); i++ {
		if lists[i] == nil {
			continue
		}

		if lists[i].Val < minValue {
			minValue = lists[i].Val
			minNdx = i
		}
	}

	if minNdx == -1 {
		return nil
	}

	lists[minNdx] = lists[minNdx].Next

	return &ListNode{minValue, nil}
}

func mergeKLists(lists []*ListNode) *ListNode {
	root := nextNode(lists)
	prev := root
	for prev != nil {
		next := nextNode(lists)
		prev.Next = next
		prev = next
	}

	return root
}
// Given a circular linked list, implement an algorithm which returns node at the begin- ning of the loop
// DEFINITION
// Circular linked list: A (corrupt) linked list in which a node’s next pointer points to an earlier node, so as to make a loop in the linked list
// EXAMPLE
// input: A -> B -> C -> D -> E -> C [the same C as earlier]
// output: C
package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func main() {
	ar := []int { 3, 1, 5, 8, 7, 5 }
	root := generateCorruptedList(ar)

	cycleStart := findNodeInCycle(root)
	fmt.Printf("First cycle node is %v \n", cycleStart.Value)
}

func findNodeInCycle(root *Node) *Node {
	m := make(map[int]bool)
	curNode := root
	for curNode != nil {
		if m[curNode.Value] == true {
			return curNode
		}

		m[curNode.Value] = true
		curNode = curNode.Next
	}

	return nil
}

func generateCorruptedList(ar []int) *Node {
	var prev, root *Node = nil, nil
	m := make(map[int]*Node)
	for _, v := range ar {
		node := Node { v, nil }
		if (m[v] == nil) {
			m[v] = &node
		} else {
			node = *m[v]
		}

		if prev != nil {
			prev.Next = &node
		} else {
			root = &node
		}

		prev = &node
	}

	return root
}

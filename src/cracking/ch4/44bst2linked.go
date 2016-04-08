// Given a binary search tree, design an algorithm which creates a linked list of all the
// nodes at each depth (i.e., if you have a tree with depth D, you’ll have D linked lists).
package main

import (
	"fmt"
	"cracking/ch4/basics"
)

type LinkedNode struct {
	Value int
	Next *LinkedNode
}

func main() {
	tree := basics.GenerateNonBalancedTree()
	lists := createLinkedListsByLevel(tree)

	for i, l := range lists {
		fmt.Printf("List for level %v \n", i);
		for l != nil {
			fmt.Printf(" %v -> ", l.Value)
			l = l.Next
		}
		fmt.Printf("\n");
	}
}

func createLinkedListsByLevel(tree *basics.TreeNode) []*LinkedNode {
	result := make([]*LinkedNode, 0)
	last := make([]*LinkedNode, 0)

	var traverse func(node *basics.TreeNode, level int)
	traverse = func(node *basics.TreeNode, level int) {
		if node == nil {
			return
		}

		curNode := &LinkedNode{node.Value, nil}
		if len(result) < level {
			result = append(result, curNode)
			last = append(last, curNode)
		} else {
			previousLast := last[level - 1]
			previousLast.Next = curNode
			last[level - 1] = curNode
		}

		traverse(node.Left, level + 1)
		traverse(node.Right, level + 1)
	}

	traverse(tree, 1)

	return result
}
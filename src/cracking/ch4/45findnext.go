package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func main() {
	tree := basics.GenerateNonBalancedPTree()

	fmt.Printf("Next node to %v is %v \n", tree.Value, getNext(tree).Value)

	n7 := tree.Left.Right
	fmt.Printf("Next node to %v is %v \n", n7.Value, getNext(n7).Value)

	n3 := tree.Left.Left.Left
	fmt.Printf("Next node to %v is %v \n", n3.Value, getNext(n3).Value)
}

func getNext(node *basics.PTreeNode) *basics.PTreeNode {
	if node.Right != nil {
		result := node.Right
		for (result.Left != nil) {
			result = result.Left
		}

		return result
	}

	if node.Parent != nil {
		if node.Parent.Left == node {
			return node.Parent
		}

		return node.Parent.Parent
	}

	return nil
}

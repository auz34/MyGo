/*Design an algorithm and write code to find the first common ancestor of two nodes
in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not
necessarily a binary search tree.*/
package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func main() {
	tree := basics.GenerateNonBalancedTree()

	n8 := tree.Left.Left.Left
	n7 := tree.Right
	c78, _, _ := findCommonAncestor(tree, n7, n8)
	fmt.Printf("Common ancestor for #%v and #%v is #%v \n", n7.Value, n8.Value, c78.Value)

	n10 := tree.Left.Right
	c108, _, _ := findCommonAncestor(tree, n10, n8)
	fmt.Printf("Common ancestor for #%v and #%v is #%v \n", n10.Value, n8.Value, c108.Value)

	n9 := tree.Left.Left
	c89, _, _ := findCommonAncestor(tree, n8, n9)
	fmt.Printf("Common ancestor for #%v and #%v is #%v \n", n8.Value, n9.Value, c89.Value)
}

func findCommonAncestor(tree, node1, node2 *basics.TreeNode) (*basics.TreeNode, bool, bool) {
	if (tree == nil) {
		return nil, false, false
	}

	lres, l1, l2 := findCommonAncestor(tree.Left, node1, node2)
	rres, r1, r2 := findCommonAncestor(tree.Right, node1, node2)

	if lres != nil {
		return lres, true, true
	}

	if rres != nil {
		return rres, true, true
	}

	b1 := l1 || r1 || tree == node1
	b2 := l2 || r2 || tree == node2

	if b1 && b2 {
		return tree, b1, b2
	}

	return nil, b1, b2
}
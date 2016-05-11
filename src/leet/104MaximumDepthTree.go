/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
 */

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node15 := &TreeNode{15, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node20 := &TreeNode{20, node15, node7}
	node3 := &TreeNode{3, node9, node20}

	fmt.Printf("%v", maxDepth(node3))
}

func max104(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max104(maxDepth(root.Left), maxDepth(root.Right))
}

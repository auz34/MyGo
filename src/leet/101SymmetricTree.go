/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
*/
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node3a := &TreeNode{3, nil, nil}
	node3b := &TreeNode{3, nil, nil}

	node4a := &TreeNode{4, nil, nil}
	node4b := &TreeNode{4, nil, nil}

	node2a := &TreeNode{ 2, node3a, node4a }
	node2b := &TreeNode{ 2, node4b, node3b }

	node1 := &TreeNode{ 1, node2a, node2b }

	fmt.Printf("%v\n", isSymmetric(node1))
}

func areMirrors(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && areMirrors(node1.Left, node2.Right) && areMirrors(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return areMirrors(root.Left, root.Right)
}
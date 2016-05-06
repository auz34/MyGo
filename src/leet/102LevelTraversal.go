/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree {3,9,20,#,#,15,7},
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
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

	result := levelOrder(node3)
	for i:=0; i<len(result); i++ {
		fmt.Printf("%v\n", result[i])
	}
}

type treeByLevels struct {
	values [][]int
}

func collectLevels(node *TreeNode, level int, result *treeByLevels) {
	if node == nil {
		return
	}

	if !(level<len(result.values)) {
		result.values = append(result.values, []int{ node.Val})
	} else {
		result.values[level] = append(result.values[level], node.Val)
	}

	collectLevels(node.Left, level+1, result)
	collectLevels(node.Right, level+1, result)
}

func levelOrder(root *TreeNode) [][]int {
	result := &treeByLevels{make([][]int, 0)}
	collectLevels(root, 0, result)
	return result.values
}
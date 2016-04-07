package basics

import "fmt"

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func GenerateNonBalancedTree() *TreeNode {
	/*    5
		 / \
		6   7
	   / \
	  9   10
	 /
	8*/
	n1 := &TreeNode{8, nil, nil}
	n1 = &TreeNode{9, n1, nil}
	n2 := &TreeNode{10, nil, nil}

	n1 = &TreeNode{6, n1, n2}
	n2 = &TreeNode{7, nil, nil}

	root := &TreeNode{5, n1, n2}


	return root
}

func traverseAndCollect(node *TreeNode, level int, results map[int]string) {
	if (node == nil) {
		return
	}

	if _, ok := results[level]; !ok {
		results[level] = ""
	}

	left := "nil"
	if node.Left != nil {
		left = fmt.Sprintf("%v", node.Left.Value)
	}

	right := "nil"
	if node.Right != nil {
		right = fmt.Sprintf("%v", node.Right.Value)
	}
	results[level] += fmt.Sprintf(" %v (l:%v; r:%v)  ", node.Value, left, right)
	traverseAndCollect(node.Left, level + 1, results)
	traverseAndCollect(node.Right, level + 1, results)
}

func PrintTreeByLevels(tree *TreeNode) {
	results := make(map[int]string)
	traverseAndCollect(tree, 0, results)

	for level, value := range results {
		fmt.Printf("%v: %s \n", level, value)
	}
}

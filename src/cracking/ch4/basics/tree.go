package basics

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

package basics

type PTreeNode struct {
	Value int
	Parent *PTreeNode
	Left *PTreeNode
	Right *PTreeNode
}

func GenerateNonBalancedPTree() *TreeNode {
	/*    5
		 / \
		6   7
	   / \
	  9   10
	 /
	8*/
	n5 := &PTreeNode{8, nil, nil, nil}
	n6 := &PTreeNode{9, nil, nil, nil}
	n7 := &PTreeNode{10, nil, nil, nil}
	n8 := &PTreeNode{8, nil, nil, nil}
	n9 := &PTreeNode{9, nil, nil, nil}
	n10 := &PTreeNode{10, nil, nil, nil}

	n5.Left = n6; n5.Right = n7
	n6.Parent = n5; n6.Left = n9; n6.Right = n10
	n7.Parent = n5
	n8.Parent = n9
	n9.Parent = n6; n9.Left = n8
	n10.Parent = n6

	return n5
}

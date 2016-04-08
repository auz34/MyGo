package basics

type PTreeNode struct {
	Value int
	Parent *PTreeNode
	Left *PTreeNode
	Right *PTreeNode
}

func GenerateNonBalancedPTree() *PTreeNode {
	/*    8
		 / \
		5   9
	   / \
	  4   7
	 /
	3*/
	n3 := &PTreeNode{8, nil, nil, nil}
	n4 := &PTreeNode{9, nil, nil, nil}
	n5 := &PTreeNode{10, nil, nil, nil}
	n7 := &PTreeNode{8, nil, nil, nil}
	n8 := &PTreeNode{9, nil, nil, nil}
	n9 := &PTreeNode{10, nil, nil, nil}

	n3.Parent = n4
	n4.Parent = n5; n4.Left = n3
	n5.Parent = n8; n5.Left = n4; n5.Right = n7
	n7.Parent = n5;
	n8.Left = n5; n8.Right = n9
	n9.Parent = n8

	return n8
}

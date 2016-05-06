/*
Given a binary tree where all the right nodes are either leaf nodes with a sibling
(a left node that shares the same parent node) or empty, flip it upside down and turn it into a tree
where the original right nodes turned into left leaf nodes. Return the new root.

For example:
Given a binary tree {1,2,3,4,5},
    1
   / \
  2   3
 / \
4   5
return the root of the binary tree [4,5,2,#,#,3,1].
   4
  / \
 5   2
    / \
   3   1
confused what "{1,#,2,3}" means?
*/

package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

type queue struct {
	internal []*TreeNode
}

func newQueue() *queue {
	return &queue{make([]*TreeNode,0)}
}

func (self *queue) isEmpty() bool {
	return len(self.internal) == 0
}

func (self *queue) push(node *TreeNode) {
	self.internal = append(self.internal, node)
}

func (self *queue) dequeue() *TreeNode {
	result := self.internal[0]
	self.internal = self.internal[1:]
	return result
}

func main() {
	node6 := &TreeNode{6, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{4, node6, node7}
	node5 := &TreeNode{5, nil, nil}
	node2 := &TreeNode{2, node4, node5}
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, node2, node3}

	fmt.Print("Source: \n")
	printTree(node1)
	fmt.Print("Result: \n")
	printTree(upsideDownBinaryTree(node1))
}

func printTree(root *TreeNode) {
	q := newQueue()
	q.push(root)

	for !q.isEmpty() {
		node := q.dequeue()

		if node != nil {
			fmt.Printf(" %v ", node.Val)
			q.push(node.Left)
			q.push(node.Right)
		} else {
			fmt.Print(" # ")
		}
	}
	fmt.Print("\n")
}

func rotateTree(node, parentOfNode *TreeNode) (*TreeNode /*where to append*/, *TreeNode /*parent*/) {
	if (node.Left == nil) {
		res := &TreeNode{node.Val, nil, nil}
		return res, res
	}

	appendTo, globalParent := rotateTree(node.Left, node)

	appendTo.Left = node.Right
	appendTo.Right = &TreeNode{node.Val, nil, nil}

	return appendTo.Right, globalParent
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	_, res := rotateTree(root, nil)
	return res
}

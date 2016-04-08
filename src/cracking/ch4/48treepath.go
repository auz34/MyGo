/*You are given a binary tree in which each node contains a value. Design an algorithm
to print all paths which sum up to that value. Note that it can be any path in the tree
- it does not have to start at the root.*/
package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func main() {
	tree := basics.GenerateNonBalancedTree2()
	fmt.Printf("Pathes for 26: \n")
	findPath(tree, 26, 26, "")
}

func findPath(tree *basics.TreeNode, total, sumLeft int, path string) {
	if tree == nil {
		return
	}

	curpath := fmt.Sprintf("%v -> %v", path, tree.Value)
	if tree.Value == sumLeft {
		fmt.Printf("Solution: %v \n", curpath)
	}

	sumLeft = sumLeft - tree.Value
	if sumLeft > 0 {
		findPath(tree.Left, total, sumLeft, curpath)
		findPath(tree.Right, total, sumLeft, curpath)
	}

	findPath(tree.Left, total, total, "")
	findPath(tree.Right, total, total, "")
}
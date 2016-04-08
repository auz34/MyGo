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
	findPath(tree, 26, make([]candidate, 0))
}

type candidate struct {
	path string
	sumLeft int
}

func findPath(tree *basics.TreeNode, total int, candidates []candidate) {
	if tree == nil {
		return
	}

	if tree.Value == total {
		fmt.Printf("Single element solution: %v \n", tree.Value)
	}

	newCandidates := make([]candidate, 0)
	for _, c := range candidates {
		if c.sumLeft == tree.Value {
			fmt.Printf("Solution: %v -> %v \n", c.path, tree.Value)
		} else if c.sumLeft > tree.Value {
			newPartialPath := fmt.Sprintf("%v -> %v", c.path, tree.Value)
			newSumLeft := c.sumLeft - tree.Value
			newCand := candidate{newPartialPath, newSumLeft}
			newCandidates = append(newCandidates, newCand)
		}
	}

	newCandidates = append(newCandidates, candidate{fmt.Sprintf("%v", tree.Value), total - tree.Value})

	findPath(tree.Left, total, newCandidates)
	findPath(tree.Right, total, newCandidates)
}
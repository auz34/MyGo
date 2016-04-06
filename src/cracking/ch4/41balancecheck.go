/*Implement a function to check if a tree is balanced. For the purposes of this question,
a balanced tree is defined to be a tree such that no two leaf nodes differ in distance
from the root by more than one.*/
package main

import (
	"fmt"
	"math"
	"cracking/ch4/basics"
)

func isBalanced(tree *basics.TreeNode) (bool, int) {
	if tree == nil {
		return true, 0
	}

	b1, h1 := isBalanced(tree.Left)
	b2, h2 := isBalanced(tree.Right)

	b := b1 && b2 && (math.Abs(float64(h1-h2)) < 2)
	h :=  1 + int(math.Max(float64(h1), float64(h2)))

	return b, h
}

func main() {
	tree := basics.GenerateNonBalancedTree()

	balanced, _ := isBalanced(tree)
	fmt.Printf("Tree is balanced: %v \n", balanced)
}


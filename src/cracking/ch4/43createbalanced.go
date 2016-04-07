// Given a sorted (increasing order) array, write an algorithm to create a binary tree with
// minimal height.

package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func main() {
	fmt.Print("First Tree \n")
	ar := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	tree := buildTree(ar)
	basics.PrintTreeByLevels(tree)
	fmt.Print("Second Tree \n")
	ar = []int {1, 2, 3, 4, 5, 6, 7 }
	tree = buildTree(ar)
	basics.PrintTreeByLevels(tree)
}

func buildTree(ar []int) *basics.TreeNode {
	l := len(ar)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return &basics.TreeNode{ar[0], nil, nil }
	}

	mid := int(l/2)
	result := &basics.TreeNode{ar[mid], nil, nil }

	if mid > 0 {
		result.Left = buildTree(ar[0:mid])
	}

	if mid < l - 1 {
		result.Right = buildTree(ar[mid+1: l])
	}

	return result
}



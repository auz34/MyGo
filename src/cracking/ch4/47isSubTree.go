/*You have two very large binary trees: T1, with millions of nodes, and T2, with hundreds
of nodes. Create an algorithm to decide if T2 is a subtree of T1.*/
package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func main() {
	T1 := basics.GenerateNonBalancedTree2()
	T2 := basics.GenerateNonBalancedTree2().Left.Right

	fmt.Printf("first test pair is subtree: %v \n", isSubtree(T1, T2))
	F2 := basics.GenerateNonBalancedTree()

	fmt.Printf("second test pair is subtree: %v \n", isSubtree(T1, F2))

}

func matchTree(t1, t2 *basics.TreeNode) bool {
	if (t1 == nil && t2 == nil) {
		return true
	}

	if (t1 == nil || t2 == nil) {
		return false
	}

	if (t1.Value != t2.Value) {
		return false
	}

	return matchTree(t1.Left, t2.Left) && matchTree(t1.Right, t2.Right)
}

func isSubtree(t1, t2 *basics.TreeNode) bool {
	if (t1 == nil) {
		return false
	}

	if matchTree(t1, t2) {
		return true
	}

	return isSubtree(t1.Left, t2) || isSubtree(t1.Right, t2)
}

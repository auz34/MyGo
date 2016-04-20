/* Imagine you are reading in a stream of integers. Periodically, you wish to be able to look up the rank of a number x
(the number of values less than or equal x), Implement the data structures and algorithms to support these operations.
That is, implement the method track(int x), which is called when each number is generated, and the
method getRankOfNumber(int x), which returns the number of values less than or equal to x (not including x itself)
EXAMPLE
Stream (in order of appearance): 5, 1, 4, 4, 5, 9, 7, 13, 3
getRankOfNumber(1) = 0
getRankOfNumber(3) = 1
getRankOfNumber(4) = 3
*/

package main

import (
	"fmt"
	"cracking/ch10/rbt"
)

func main() {
	tree := rbt.NewRedBlackTree()

	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(3)

	fmt.Printf("Rank of 1 is: %v \n", tree.GetRank(1))
	fmt.Printf("Rank of 3 is: %v \n", tree.GetRank(3))
	fmt.Printf("Rank of 4 is: %v \n", tree.GetRank(4))
	fmt.Printf("Rank of 9 is: %v \n", tree.GetRank(9))
}
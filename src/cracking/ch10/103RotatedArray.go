/*Given a sorted array of n integers that has been rotated an unknown number of times, give an O(log n)
algorithm that finds an element in the array. You may assume that the array was originally sorted in increasing order.
EXAMPLE:
Input:  nd 5 in array (15 16 19 20 25 1 3 4 5 7 10 14)
Output: 8 (the index of 5 in the array)*/
package main

import "fmt"


func main() {
	ar := []int {15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}

	fmt.Printf("search for 15 find it at %v \n", search(ar, 15, 0, len(ar) - 1))
	fmt.Printf("search for 19 find it at %v \n", search(ar, 19, 0, len(ar) - 1))
	fmt.Printf("search for 25 find it at %v \n", search(ar, 25, 0, len(ar) - 1))
	fmt.Printf("search for 7 find it at %v \n", search(ar, 7, 0, len(ar) - 1))
	fmt.Printf("search for 14 find it at %v \n", search(ar, 14, 0, len(ar) - 1))
	fmt.Printf("search for 13 find it at %v \n", search(ar, 13, 0, len(ar) - 1))
}

func search(list []int, num, left, right int) int {
	if left > right {
		return -1
	}

	mid := (right + left) / 2
	if list[mid] == num {
		return mid
	}

	result := -1
	if list[mid] < num {
		result = search(list, num, mid+1, right)
	} else {
		result = search(list, num, left, mid - 1)
	}

	if result == -1 && list[left] > list[right] {
		if list[mid] < num {
			result = search(list, num, left, mid - 1)
		} else {
			result = search(list, num, mid+1, right)
		}
	}

	return result
}
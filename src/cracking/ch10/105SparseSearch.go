/*
Given a sorted array of strings which is interspersed with empty strings, write a method to find
the location of a given string.
Example: find “ball” in [“at”, “”, “”, “”, “ball”, “”, “”, “car”, “”, “”, “dad”, “”, “”] will return 4
Example:  nd “ballcar” in [“at”, “”, “”, “”, “”, “ball”, “car”, “”, “”, “dad”, “”, “”] will return -1
*/

package main

import "fmt"

func main() {
	ar1 := []string {"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", "" }
	ar2 := []string {"at", "", "", "", "", "ball", "car", "", "", "dad", "", "" }

	fmt.Printf("Search for ball: %v \n", sparseSearch(ar1, "ball", 0, len(ar1)))
	fmt.Printf("Search for carball: %v \n", sparseSearch(ar2, "carball", 0, len(ar2)))
}

func intMin(i1, i2 int) int {
	if i1<i2 {
		return i1
	}

	return i2
}

func intMax(i1, i2 int) int {
	if i1<i2 {
		return i2
	}

	return i1
}

func sparseSearch(list []string, key string, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	realMid := mid

	for list[realMid] == "" && realMid >= left {
		realMid--
	}

	if list[realMid] == "" {
		realMid = mid
		for list[realMid] == "" && realMid <= right {
			realMid++
		}
	}

	if list[realMid] == "" {
		return -1
	}

	if list[realMid] == key {
		return realMid
	}

	if list[realMid] > key {
		return sparseSearch(list, key, left, intMin(mid, realMid) - 1)
	}

	return sparseSearch(list, key, intMax(mid, realMid) + 1, right)
}
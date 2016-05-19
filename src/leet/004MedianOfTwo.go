/*
There are two sorted arrays nums1 and nums2 of size m and n respectively. Find the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).
*/
package main

import "fmt"

func main() {
	ar1 := []int { 1, 2, 3, 4, 5 }
	ar2 := []int { 6, 7, 8, 9, 10 }
	fmt.Printf("Answer #1: %v\n", findMedianSortedArrays(ar1, ar2))
}

func min004(x, y int) int {
	if x<y {
		return x
	}

	return y
}

func max004(x, y int) int {
	if x>y {
		return x
	}

	return y
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 2 && len(nums2) == 2 {
		// a, b - c, d
		return float64(max004(nums1[0], nums2[0]) + min004(nums1[1], nums2[1])) / 2
	}

	mid1 := len(nums1) / 2
	mid2 := len(nums2) / 2

	if nums1[mid1] == nums2[mid2] {
		return float64(nums1[mid1])
	}

	if nums1[mid1] <nums2[mid2] {
		return findMedianSortedArrays(nums1[mid1:], nums2[:mid2+1])
	}

	return findMedianSortedArrays(nums1[:mid1+1], nums2[mid2:])

}

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

func singleMedian(nums []int) float64{
	if len(nums) == 0 {
		return -1
	}

	if len(nums) % 2 != 0 {
		return nums[len(nums)/2]
	}

	mid := len(nums) /2
	return (nums[mid] + nums[mid-1]) / 2
}

func findMedian4SpecialCases(nums1, nums2 []int) float64 {
	d1 := len(nums2)
	if d1>2 {
		d1 = 3
	}
	caseNum := 10*d1 + len(nums1)

	switch (caseNum) {
	case 0: // |nums2| = 0 |nums1| = 0
		return float64(-1)
	case 1: // |nums2| = 0 |nums1| = 1
		return float64(nums1[0])
	case 2: // |nums2| = 0 |nums1| = 2
		return float64(nums1[0] + nums1[1])/2
	case 10: //|nums2| = 1 |nums1| = 0
		return float64(nums2[0])
	case 11: //|nums2| = 1 |nums1| = 1
		return float64(nums1[0] + nums2[0])/2
	case 12: //|nums2| = 1 |nums1| = 2
		if nums2[0]<nums1[0] {
			return float64(nums1[0])
		}
		return float64(min004(nums2[0], nums1[1]))
	case 21: //|nums2| = 2 |nums1| = 1
		if nums1[0]<nums2[0] {
			return float64(nums2[0])
		}
		return float64(min004(nums1[0], nums2[1]))
	case 22: //|nums2| = 2 |nums1| = 1
		return float64(max004(nums1[0], nums2[0]) + min004(nums1[1], nums2[1])) / 2
	case 30: //|nums2| >= 3 |nums1| = 0
		return singleMedian(nums2)
	case 31:
	case 32:
		panic("Have not implemented yet")
	}

	if len(nums1) == 0 {
		return singleMedian(nums2)
	}

	if len(nums2) == 0 {
		return singleMedian(nums1)
	}

	if len(nums1) == 1 {

	}


	if len(nums1) == 2 && len(nums2) == 2 {
		// a, b - c, d
		return float64(max004(nums1[0], nums2[0]) + min004(nums1[1], nums2[1])) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) <= 2 {
		return findMedian4SpecialCases(nums1, nums2)
	}

	if len(nums2) <= 2 {
		return findMedian4SpecialCases(nums2, nums1)
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

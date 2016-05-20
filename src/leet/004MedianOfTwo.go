/*
There are two sorted arrays nums1 and nums2 of size m and n respectively. Find the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).
*/
package main

import "fmt"

func main() {
	fmt.Print("Test special cases: ")

	ar1 := []int {  }
	ar2 := []int {  }
	fmt.Printf("Case #1:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 5 }
	ar2 = []int {  }
	fmt.Printf("Case #2:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 5 }
	ar2 = []int { 1 }
	fmt.Printf("Case #3:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 5 }
	ar2 = []int {  3 }
	fmt.Printf("Case #4:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 3, 9 }
	ar2 = []int {  1 }
	fmt.Printf("Case #5:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 3, 9 }
	ar2 = []int {  15 }
	fmt.Printf("Case #6:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 3, 9 }
	ar2 = []int { 5, 7 }
	fmt.Printf("Case #7:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 3, 9 }
	ar2 = []int { 5, 11 }
	fmt.Printf("Case #8:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 3, 9 }
	ar2 = []int { 11, 13 }
	fmt.Printf("Case #9:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #10:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 6 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #11:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #12:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 10 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #13:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1 }
	ar2 = []int { 5, 7, 9, 11 }
	fmt.Printf("Case #14:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 6 }
	ar2 = []int { 5, 7, 9, 11 }
	fmt.Printf("Case #15:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8 }
	ar2 = []int { 5, 7, 9, 11 }
	fmt.Printf("Case #16:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 10 }
	ar2 = []int { 5, 7, 9, 11 }
	fmt.Printf("Case #17:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 2 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #18:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 6 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #19:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 8 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #20:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8, 10 }
	ar2 = []int { 5, 7, 9 }
	fmt.Printf("Case #21:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8, 10 }
	ar2 = []int { 5, 7, 11 }
	fmt.Printf("Case #22:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 12, 13 }
	ar2 = []int { 5, 7, 11 }
	fmt.Printf("Case #23:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 2 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #24:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 5 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #25:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 10 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #26:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 12 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #27:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 16 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #28:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8, 9 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #29:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 8, 12 }
	ar2 = []int { 3, 7, 11, 15 }
	fmt.Printf("Case #30:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	fmt.Print("Test real cases: ")

	ar1 = []int { 1, 2, 3, 4, 5 }
	ar2 = []int { 6, 7, 8 }
	fmt.Printf("Case #31:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 2, 3, 4, 5 }
	ar2 = []int { 6, 7, 8, 9 }
	fmt.Printf("Case #32:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 2, 3, 4, 5 }
	ar2 = []int { 6, 7, 8, 9, 10 }
	fmt.Printf("Case #33:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))

	ar1 = []int { 1, 3, 5, 7, 9 }
	ar2 = []int { 2, 4, 6 }
	fmt.Printf("Case #34:\n nums1=%v\n nums2=%v\n answer=%v\n", ar1, ar2, findMedianSortedArrays(ar1, ar2))
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
		return float64(nums[len(nums)/2])
	}

	mid := len(nums) /2
	return float64(nums[mid] + nums[mid-1]) / 2
}

func medianArrayAndNewElement(nums []int, el int) float64 {
	mid := len(nums) / 2
	if (len(nums) % 2 != 0) {
		if el<nums[mid] {
			if el<nums[mid-1] {
				return float64(nums[mid] + nums[mid-1]) / 2
			}

			return float64(nums[mid] + el) / 2
		}

		if el>nums[mid+1] {
			return float64(nums[mid] + nums[mid+1]) / 2
		}

		return float64(nums[mid] + el) / 2
	}

	if el<nums[mid] {
		if el<nums[mid-1] {
			return float64(nums[mid-1])
		}

		return float64(el)
	}

	return float64(nums[mid])
}

func medianArrayAndTwoNewElement(nums []int, el1, el2 int) float64 {
	mid := len(nums) / 2
	if len(nums) % 2 != 0 {
		if el1 < nums[mid]  {
			if el2 > nums[mid] {
				return float64(nums[mid])
			}

			return float64(max004(el2, nums[mid-1]))
		}

		return float64(min004(el1, nums[mid+1]))
	}

	if el1 < nums[mid-1]  {
		if el2 > nums[mid] { // el1 < nums[mid-1] && el2 > nums[mid]
			return float64(nums[mid-1] + nums[mid]) / 2
		}

		if el2 < nums[mid-1] {
			return float64(nums[mid-1] + max004(el2, nums[mid-2])) / 2
		}

		return float64(nums[mid-1] + el2) / 2
	}

	if el1 < nums[mid] {
		return float64(el1 + min004(el2, nums[mid])) / 2
	}

	return float64(nums[mid] + nums[mid+1]) / 2
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
		return medianArrayAndNewElement(nums2, nums1[0])
	case 32:
		return medianArrayAndTwoNewElement(nums2, nums1[0], nums1[1])
	default:
		panic("not a special case")
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) <= 2 {
		return findMedian4SpecialCases(nums1, nums2)
	}

	if len(nums2) <= 2 {
		return findMedian4SpecialCases(nums2, nums1)
	}

	if (len(nums1) > len(nums2)) {
		return findMedianSortedArrays(nums2, nums1)
	}

	mid1 := len(nums1) / 2
	mid2 := len(nums2) / 2

	if nums1[mid1] == nums2[mid2] {
		return float64(nums1[mid1])
	}

	if nums1[mid1] <nums2[mid2] {
		if len(nums1) != len(nums2) {
			return findMedianSortedArrays(nums1[mid1:], nums2[:mid2])
		}

		return findMedianSortedArrays(nums1[mid1:], nums2[:mid2+1])

	}

	if len(nums1) != len(nums2) {
		return findMedianSortedArrays(nums1[:mid1 + 1], nums2[mid2 - 1:])
	}

	return findMedianSortedArrays(nums1[:mid1 + 1], nums2[mid2:])

}

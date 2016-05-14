/*
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:
Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].

Example 2:
Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].

This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
 */

package main

import "fmt"

func main() {
	set1 := []Interval { Interval{1,3}, Interval{6,9}}
	fmt.Print("Input #1:\n")
	printIntervals(set1)
	fmt.Print("Output #1:\n")
	printIntervals(insert(set1, Interval{2,5}))

	set2 := []Interval { Interval{1,2}, Interval{3,5}, Interval{6,7}, Interval{8,10}, Interval{12,16}}
	fmt.Print("Input #1:\n")
	printIntervals(set2)
	fmt.Print("Output #1:\n")
	printIntervals(insert(set2, Interval{4,9}))
}

func printIntervals(intervals []Interval) {
	for i:=0; i<len(intervals); i++ {
		fmt.Printf(" [%v, %v] ", intervals[i].Start, intervals[i].End)
	}

	fmt.Print("\n")
}

type Interval struct {
	Start int
	End   int
}

func min57(x,y int) int {
	if x<y {
		return x
	}

	return y
}

func max57(x,y int) int {
	if x>y {
		return x
	}

	return y
}

func merge57(int1, int2 Interval) Interval {
	return Interval{min57(int1.Start, int2.Start), max57(int1.End, int2.End)}
}

func compare57(int1, int2 Interval) int {
	if int1.End < int2.Start {
		return -1
	}

	if int2.End < int1.Start {
		return 1
	}

	return 0
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	result := make([]Interval, 0)

	for i:=0; i<len(intervals); i++ {
		cmp := compare57(intervals[i], newInterval)

		if cmp < 0 {
			result = append(result, intervals[i])
		} else if cmp == 0 {
			newInterval = merge57(intervals[i], newInterval)
		} else {
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		}
	}

	result = append(result, newInterval)

	return result
}
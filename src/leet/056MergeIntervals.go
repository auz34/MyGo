/*Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].*/


package main

import "fmt"

func main() {
	intervals := []Interval{
		Interval{1,3},
		Interval{2,6},
		Interval{8,10},
		Interval{15,18},
	}

	print(merge(intervals))

}

func print(intervals []Interval) {
	for i:=0; i<len(intervals); i++ {
		fmt.Printf("[%v,%v],", intervals[i].Start, intervals[i].End)
	}
	fmt.Print("\n")
}


//Definition for an interval.
type Interval struct {
   Start int
   End   int
}

func isOverlap(a,b Interval) bool {
   return compare(a, b) == 0
}

func compare(a, b Interval) int {
   if b.End < a.Start {
      return a.Start - b.End
   }

   if a.End < b.Start {
      return a.End - b.Start
   }

   return 0
}

func min(x,y int) int {
   if x < y {
      return x
   }

   return y
}

func max(x,y int) int {
   if x > y {
      return x
   }

   return y
}

func mergeIntervals(target, source *Interval) {
	if !isOverlap(*target, *source) {
		panic("Wrong logic")
	}

	target.Start = min(target.Start, source.Start)
	target.End = max(target.End, source.End)
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
	   return intervals
	}

	left := make([]Interval, 0)
	right := make([]Interval, 0)

	mid := len(intervals) / 2
	intervals[0], intervals[mid] = intervals[mid], intervals[0]
	v := intervals[0]
	overlaps := false

	for i:=1; i<len(intervals); i++ {
		cmp := compare(v, intervals[i])
		if cmp < 0 {
			right = append(right, intervals[i])
		} else if cmp>0 {
			left = append(left, intervals[i])
		} else {
			overlaps = true
			mergeIntervals(&v, &intervals[i])
		}
	}

	if overlaps {
		modifiedIntervals := append(append(left, v), right...)
		return merge(modifiedIntervals)
	} else {
		return append(append(merge(left), v), merge(right)...)
	}
}

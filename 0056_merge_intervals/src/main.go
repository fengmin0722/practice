package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals);{
		start := intervals[i][0]
		end := intervals[i][1]

		for {
			if i < len(intervals) && intervals[i][0] <= end{
				if intervals[i][1] > end {
					end = intervals[i][1]
				}
				i++
				continue
			}
			tmp := make([]int,2)
			tmp[0] = start
			tmp[1] = end
			ret = append(ret, tmp)
			break
		}
	}
	return ret
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18},}
	fmt.Println(mergeIntervals(intervals))
}

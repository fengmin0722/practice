package main

import (
	"fmt"
	"math"
)

func isStraight(nums []int) bool {
	min, max := 14, 0
	mark := make(map[int]bool)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if mark[num] == true {
			return false
		}

		mark[num] = true
		min = int(math.Min(float64(min), float64(num)))
		max = int(math.Max(float64(max), float64(num)))
	}

	if max - min + 1 <= 5 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isStraight([]int{1,2,3,4,5}))
	fmt.Println(isStraight([]int{0,0,1,2,5}))
}

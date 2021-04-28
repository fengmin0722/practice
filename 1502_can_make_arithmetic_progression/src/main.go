package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(nums []int) bool {
	sort.Ints(nums)
	delta := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i-1] != delta {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
}

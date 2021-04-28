package main

import "fmt"

func searchInsertPos(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + (high - low) / 2
		switch {
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		default:
			return mid
		}
	}
	return low
}

func main() {
	nums := []int{2, 3, 5, 6}
	target := 1
	fmt.Println(searchInsertPos(nums, target))
}

package main

import "fmt"

func guessNumber_r(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + (high - low) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return guessNumber_r(nums, mid + 1, high, target)
	} else {
		return guessNumber_r(nums, low, mid - 1, target)
	}
}

func guess(num, target int) int {
	if num == target {
		return 0
	} else if num < target {
		return -1
	} else {
		return 1
	}
}

func guessNumber(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + (high - low) / 2
		ret := guess(nums[mid], target)
		if ret == 0 {
			return mid
		} else if ret == -1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(guessNumber(nums, 5))
}

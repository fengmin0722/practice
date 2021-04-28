package main

import (
	"fmt"
	"reflect"
)

// 递归实现
func binarySearch_r(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearch_r(nums, mid + 1, high, target)
	} else {
		return binarySearch_r(nums, low, mid - 1, target)
	}
}

// 非递归实现
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(reflect.TypeOf(nums))
	fmt.Println(binarySearch(nums, 7))
	fmt.Println(binarySearch(nums, 100))
}

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	l := searchRange2(nums, true, target)
	r := searchRange2(nums, false, target)
	return []int{l, r}
}

func searchRange2(nums []int, leftTest bool, target int) int {
	l , r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		switch {
		case nums[mid] < target:
			l = mid + 1
		case nums[mid] > target:
			r = mid - 1
		default:
			if leftTest {
				if mid == 0 || nums[mid] > nums[mid-1] { // 不再继续向左走的 2 个边界条件
					return mid
				}
				r = mid - 1 // 继续在左侧找
			} else {
				if mid == len(nums)-1 || nums[mid] < nums[mid+1] {
					return mid
				}
				l = mid + 1 // 继续在右侧找
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange([]int{1}, 1))                 // [0, 0]
}

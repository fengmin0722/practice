package main

import "fmt"

func tidyColor(nums []int) {
	i, j := 0, len(nums) - 1
	cur := 0
	for cur <= j {
		if nums[cur] == 0 {
			nums[cur], nums[i] = nums[i], nums[cur]
			i++
			cur++
		} else if (nums[cur] == 2) {
			nums[cur], nums[j] = nums[j], nums[cur]
			j--
		} else {
			cur++
		}
	}
}

func main() {
	nums := []int{2, 1, 2, 1, 0, 0}
	tidyColor(nums)
	fmt.Println(nums)
}

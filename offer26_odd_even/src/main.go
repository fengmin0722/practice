package main

import "fmt"

func tidyOddEven(nums []int) []int {
	if len(nums) <= 0 {
		return nums
	}

	i, j := 0, len(nums) - 1
	for i <= j {
		if nums[i] % 2 == 1 {
			i++
			continue
		}

		if nums[j] % 2 == 0 {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(tidyOddEven(nums))
}

package main

import "fmt"

// 两次遍历
func twoSum1(nums []int32, target int32) []int32 {
	res := []int32{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				res = append(res, int32(i))
				res = append(res, int32(j))
				return res
			}
		}
	}
	return res
}

// 一次遍历
func twoSum2(nums []int32, target int32) []int32 {
	res := []int32{}
	numMap := make(map[int32]int)
	for i := 0; i < len(nums); i++ {
		desNum := target - nums[i]
		if _, ok := numMap[desNum]; ok {
			res = append(res, int32(numMap[desNum]))
			res = append(res, int32(i))
			return res
		} else {
			numMap[nums[i]] = i
		}
	}
	return res
}

func main() {
	nums := []int32{2, 7, 11, 15}
	target := int32(9)

	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println("hello")
}
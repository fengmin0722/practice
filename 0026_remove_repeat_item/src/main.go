package main

import "fmt"

// 针对有序数组，双指针法是十分常见且有用的
func removeDuplicates(nums []int32) ([]int32, int) {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { // 相邻的数不相等
			slow++
			fast++
			nums[slow] = nums[fast] // 将最新的新数存储到慢指针的位置
			continue
		}
		fast++
	}
	return nums, slow + 1
}

func main() {
	array := []int32{1, 1, 2, 3, 3, 5}
	newArray, size := removeDuplicates(array)
	fmt.Println(newArray)
	fmt.Println(len(newArray))
	fmt.Println(size)
	newNewArray := newArray[:size]
	fmt.Println(newNewArray)
}

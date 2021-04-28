package main

import "fmt"

func nextGreatestLetter(ch []byte, target byte) int {
	left, right := 0, len(ch) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if ch[mid] > target {
			if mid != 0 && ch[mid - 1] <= target {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	chs := []byte{'c', 'f', 'g'}
	target := 'a'
	fmt.Println(nextGreatestLetter(chs, byte(target)))
}

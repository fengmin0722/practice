package main

import "fmt"

func findString(words []string, target string) int {
	l, r := 0, len(words) - 1
	for l <= r {
		mid := l + (r - l) / 2
		// 找出不是空字符串的mid
		if words[mid] == "" {
			left := mid - 1
			right := mid + 1
			for {
				// 没有找到
				if left < l && right > r {
					return -1
				} else if (right <= r && words[right] != "") {
					mid = right
					break
				} else if (left >= l && words[left] != "") {
					mid = left
					break
				}
				right++
				left--
			}
		}

		if words[mid] == target {
			return mid
		} else if words[mid] < target {
			l = mid + 1
		} else {
			r = mid -1
		}
	}
	return -1
}

func main() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(findString(words, "ball"))
}

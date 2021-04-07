package main

import "fmt"

func lengthOfLastWord(s string) int32 {
	// 没有去掉最后面的空格
	var count int32 = 0
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		i--
		count++
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}

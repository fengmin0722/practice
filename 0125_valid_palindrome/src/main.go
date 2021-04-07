package main

import (
	"fmt"
)

func isPalindrome(target string) bool {
	var arr []rune
	for _, ch := range target {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			arr = append(arr, ch)
		} else if ch >= 'A' && ch <= 'Z' {
			arr = append(arr, ch+('a'-'A'))
		} else {
			continue
		}
	}

	if string(arr) == reverseString(string(arr)) {
		return true
	}
	return false
}

func reverseString(target string) string {
	des := []byte(target)
	i, j := 0, len(des) - 1
	for i < j {
		des[i], des[j] = des[j], des[i]
		i++
		j--
	}
	return string(des)
}

func isPalindromeNum(x int32) bool {
	if x < 0 {
		return false
	}

	result := int32(0)
	target := x

	// 121
	for x > 0 {
		result = result * 10 + x % 10
		x /= 10
	}

	if target == result {
		return true
	}
	return false
}

func main() {
	fmt.Println(reverseString("hello"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))

	ch := 'A'
	fmt.Println(ch+('a'-'A'))

	ch1 := 'B'
	fmt.Println(ch1+('a'-'A'))

	fmt.Println(isPalindromeNum(121))
}

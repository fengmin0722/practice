package main

import "fmt"

func reverseWord(target string) string {
	des := []byte(target)
	reverseString(des)
	left, right := 0, 0
	for right < len(des) {
		if des[right] != ' ' {
			right++
			continue
		}
		reverseString(des[left:right])
		left = right + 1
		right = right + 1
	}
	reverseString(des[left:right])
	return string(des)
}

func reverseString(target []byte) {
	i, j := 0, len(target) - 1
	for i < j {
		target[i], target[j] = target[j], target[i]
		i++
		j--
	}
}

func main() {
	fmt.Println(reverseWord("i is fengmin"))
}

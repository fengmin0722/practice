package main

import "fmt"

func reverseString(target string) string {
	des := []rune(target)
	i, j := 0, len(des) - 1
	for ; i < j; {
		des[i], des[j] = des[j], des[i]
		i = i + 1
		j = j - 1
	}
	return string(des)
}

func main() {
	fmt.Println(reverseString("hello"))
}
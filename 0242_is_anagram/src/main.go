package main

import (
	"fmt"
	"sort"
)

func isAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	f := []byte(first)
	s := []byte(second)
	sort.Slice(f, func(i, j int) bool {
		return f[i] < f[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	i := 0
	for i < len(f) {
		if f[i] != s[i] {
			return false
		}
		i++
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

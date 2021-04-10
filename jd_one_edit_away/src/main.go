package main

import (
	"fmt"
	"math"
)
/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
示例 1:

输入: 
first = "pale"
second = "ple"
输出: True
 

示例 2:

输入: 
first = "pales"
second = "pal"
输出: False
 */

func oneEditAway(f []byte, s[]byte) bool {
	s1, s2 := len(f), len(s)
	if math.Abs(float64(s1 - s2)) > 1 {
		return false
	}
	l, r1, r2 := 0, s1-1, s2-1

	// 找到最左侧的位置
	for l <= r1 && l <= r2 && f[l] == s[l] {
		l++
	}

	// 找到最右侧的位置
	for r1 >= 0 && r2 >= 0 && f[r1] == f[r2] {
		r1--
		r2--
	}

	// 判断两个位置之间的距离
	if r1 - l < 1 && r2 - l < 1 {
		return true
	}
	return false
}

func main() {
	f := "pale"
	s := "pal"
	fmt.Println(oneEditAway([]byte(f), []byte(s)))
}

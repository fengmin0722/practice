package main

import (
	"fmt"
	"strings"
)

// "127" " 127"
func checkSeg(seg string) bool {
	if seg == "" {
		return false
	}

	// step 1 找出第一非空格的位置
	i := 0
	for {
		if seg[i] != ' ' {
			break
		}
		i++
	}

	// 都是空格
	if i >= len(seg) {
		return false
	}

	// step 2 求出数组和
	num := 0
	for {
		if i >= len(seg) || seg[i] == ' ' {
			break
		}
		//case '0' <=  && c <= '9':
		num = num * 10 + int(seg[i] - byte('0'))
		i++
	}

	if num < 0 || num > 255 {
		return false
	}
	fmt.Println("num : ", num)

	// step 3 处理中间有空格
	if i < len(seg) {
		for {
			// 循环跳出条件
			if i >= len(seg) {
				break
			}
			if seg[i] != ' ' {
				return false
			}
			i++
		}
	}

	return true
}

//
func checkIp(ip string) bool {
	if ip == "" {
		return false
	}

	segs := strings.Split(ip, ".")
	if len(segs) != 4 {
		return false
	}

	for i := 0; i < len(segs); i++ {
		if checkSeg(segs[i]) == false {
			return false
		}
	}
	return true
}

func main() {
	ip := "127.0.0.1"
	if checkIp(ip) {
		fmt.Println("checkIp succ, ip ", ip)
	} else {
		fmt.Println("checkIp failed, ip ", ip)
	}
	fmt.Println("hello")
}
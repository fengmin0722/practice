package main

import "fmt"

var calc map[int]int

func fab(i int) int {
	if i == 1 {
		return 1 // 只能走一步
	} else if i == 2 {
		return 2 // 走一步或者走两步
	}

	if val, ok := calc[i]; ok {
		return val
	}

	if calc == nil {
		calc = make(map[int]int)
	}

	val := fab(i-1) + fab(i-2)
	if _, ok := calc[i]; !ok {
		calc[i] = val
	}
	return calc[i]
}

func main() {
	fmt.Println(fab(3))
	fmt.Println(calc)
}

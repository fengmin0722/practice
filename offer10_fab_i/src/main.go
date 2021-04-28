package main

import "fmt"

var calc map[int]int

func fab(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	}
	if val, ok := calc[i]; ok {
		return val
	}
	val := fab(i-1) + fab(i-2)
	if calc == nil {
		calc = make(map[int]int)
	}
	calc[i] = val
	return calc[i]
}

func main() {
	fmt.Println(fab(10))
	fmt.Println(calc)
}

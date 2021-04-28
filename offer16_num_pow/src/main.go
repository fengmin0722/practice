package main

import "fmt"

var calc map[int]int
var calcex map[int]int

func pow(number, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return number
	}

	if v, ok := calc[n]; ok {
		return v
	}

	v := number * pow(number, n-1)
	if calc == nil {
		calc = make(map[int]int)
	}
	calc[n] = v
	return v
}

func multi(number, n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return number
	}

	if v, ok := calcex[n]; ok {
		return v
	}

	v := number + multi(number, n-1)
	if calcex == nil {
		calcex = make(map[int]int)
	}
	calcex[n] = v
	return calcex[n]
}

func main() {
	fmt.Println(pow(3, 3))
	fmt.Println(calc)
	fmt.Println(multi(1, 1))
	fmt.Println(calcex)
}

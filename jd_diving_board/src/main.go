package main

import (
	"fmt"
	"sort"
)

func divingBoard(shorter int, longer int, k int) []int {
	var res []int
	if k == 0 {
		return res
	}

	if shorter == longer {
		res = append(res, k * shorter)
		return res
	}

	for i := 0; i <= k; i++ {
		a := shorter * i + (k - i) * longer
		res = append(res, a)
	}

	sort.Ints(res)

	return res
}

func divingBoardEx(shorter, longer, k int) []int {
	var res []int
	for {
		if k == 0 {
			break
		}
		if shorter == longer {
			res = append(res, k * shorter)
			break
		}

		for i := 0; i <= k; i++ {
			a := shorter * i + (k - i) * longer
			res = append(res, a)
		}
		break
	}


	sort.Ints(res)
	return res
}

func main() {
	fmt.Println(divingBoard(3, 10, 2))
	fmt.Println(divingBoardEx(3, 10, 2))
}

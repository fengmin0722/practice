package main

import "fmt"

func setZeroes(matrix [][]int) {
	rows := make(map[int]int)
	columns := make(map[int]int)

	for i := range matrix {
		for j, num := range matrix[i] {
			if num == 0 {
				rows[i]++
				columns[j]++
			}
		}
	}

	// c相当于索引了，日了狗了，牛逼
	for r := range rows {
		for c := range matrix[0] {
			fmt.Println(fmt.Sprintf("r[%v], c[%v]\n", r, c))
			matrix[r][c] = 0
		}
	}

	// r相当于数组的索引了，日了狗了，牛逼
	for c := range columns {
		for r := range matrix {
			fmt.Println(fmt.Sprintf("r[%v], c[%v]\n", r, c))
			matrix[r][c] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},

		{0, 1, 2, 0},
		{3, 0, 5, 2},
		{1, 3, 1, 5},
	}

	fmt.Println(matrix)

	fmt.Println(matrix[0])

	setZeroes(matrix)

	fmt.Println(matrix)
}

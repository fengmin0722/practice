package main

import "fmt"

func mergeSortedArray(left , right []int) []int {
	array := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			array = append(array, left[i])
			i++
			continue
		}

		if left[i] > right[j] {
			array = append(array, right[j])
			j++
			continue
		}

		array = append(array, left[i])
		array = append(array, right[j])
		i++
		j++
	}

	for i < len(left) {
		array = append(array, left[i])
		i++
	}

	for j < len(right) {
		array = append(array, right[j])
		j++
	}
	return array
}

func main() {
	left := []int{1, 10, 100, 101}
	right := []int{2, 9, 99, 130, 145}
	fmt.Println(mergeSortedArray(left, right))
}

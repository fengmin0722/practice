package main

import "fmt"

func rotateLeftWords(target string, k int32) string {
	des := []byte(target)
	k = k % int32(len(target))

	reverseString(des[:k])
	reverseString(des[k:])
	reverseString(des)
	return string(des)
}

func reverseString(target []byte) {
	i, j := 0, len(target)-1
	for i < j {
		target[i], target[j] = target[j], target[i]
		i++
		j--
	}
}

func main() {
	fmt.Println("1234567 " + " Rotate 3 : " + rotateLeftWords("1234567", 3))
}

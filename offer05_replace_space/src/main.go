package main

import "fmt"

func replaceSpace(target string) string {
	var result []byte
	for _, ch := range []byte(target) {
		if ch == ' ' {
			result = append(result, '%')
			result = append(result, '2')
			result = append(result, '0')
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(replaceSpace("hello world"))
}

package main

import "fmt"

func validParentheses(target string) bool {
	s := []byte(target)
	array := make([]byte, len(s), len(s))
	top := 0
	for i := 0; i < len(s); i++ {
		switch(s[i]) {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			array[top] = s[i]
			top++
		case ')':
			top--
			if top < 0 || array[top] != '(' {
				return false
			}
		case '}':
			top--
			if top < 0 || array[top] != '{' {
				return false
			}
		case ']':
			top--
			if top < 0 || array[top] != '[' {
				return false
			}
		}
		fmt.Println(fmt.Sprintf("%c", s[i]))
	}
	if top != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(validParentheses("(){}[]"))
}

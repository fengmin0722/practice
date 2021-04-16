package main

import "fmt"

type Stack struct {
	Data []int
	Size int
	TopPos []int
}

func NewStack(size int) *Stack {
	s := &Stack{
		Data: make([]int, 3 * size, 3 * size),
		Size: size,
		TopPos: make([]int, 3, 3),
	}

	// 初始化栈顶位置
	for i := 0; i < 3; i++ {
		s.TopPos[i] = i * size
	}
	return s
}

func (s *Stack) Push(stack, val int) {
	if stack < 0 || stack >= 3 {
		return
	}
	max := (stack + 1) * s.Size
	if s.TopPos[stack] >= max {
		return
	}
	pos := s.TopPos[stack]
	s.TopPos[stack]++
	s.Data[pos] = val
}

func (s *Stack) Pop(stack int) int {
	if stack < 0 || stack >= 3 {
		return -1
	}
	if s.TopPos[stack] <= stack * s.Size {
		return -1
	}
	pos := s.TopPos[stack] - 1
	s.TopPos[stack] = s.TopPos[stack] - 1
	res := s.Data[pos]
	s.Data[pos] = 0
	return res
}

func main() {
	s := NewStack(10)
	s.Push(0, 1)
	s.Push(0, 2)
	s.Push(1, 11)
	s.Push(1, 12)
	s.Push(2, 21)
	s.Push(2, 22)
	fmt.Println(s.Data)
	fmt.Println(s.Pop(0))
	fmt.Println(s.Data)
}

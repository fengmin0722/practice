package main

import "fmt"

type Stack struct {
	elements []int
}

func NewStack(cap int) *Stack {
	return &Stack{
		elements: make([]int, 0, cap),
	}
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) Push(elem int) {
	s.elements = append(s.elements, elem)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elem
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		return -1
	}
	return s.elements[len(s.elements)-1]
}

// 用两个栈 实现队列
type Queue struct {
	in *Stack
	out *Stack
}

func NewQueue(cap int) *Queue {
	return &Queue{
		in: NewStack(cap),
		out: NewStack(cap),
	}
}

func (q *Queue) In2Out() {
	for q.in.Len() > 0 {
		q.out.Push(q.in.Pop())
	}
}

func (q *Queue) Push(elem int) {
	q.in.Push(elem)
}

func (q *Queue) Pop() int {
	if q.out.Len() <= 0 {
		q.In2Out()
	}
	return q.out.Pop()
}

func main() {
	queue := NewQueue(10)

	queue.Push(5)
	queue.Push(6)
	queue.Push(7)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	tail := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		tail.Next = node
		tail = tail.Next
	}
	return head
}

func displayList(head *ListNode) string {
	var res strings.Builder
	if head == nil {
		return res.String()
	}
	cur := head
	for cur != nil {
		res.WriteString(strconv.Itoa(cur.Val))
		res.WriteString(" ")
		cur = cur.Next
	}
	return res.String()
}

func mergeSortedLinkedList(f, s *ListNode) *ListNode {
	if f == nil && s != nil {
		return s
	}
	if f != nil && s == nil {
		return f
	}
	headRet := &ListNode{0, nil}
	tailRet := headRet
	headF := f
	headS := s
	for headF != nil && headS != nil {
		if headF.Val < headS.Val {
			tmp := headF.Next
			headF.Next = nil
			tailRet.Next = headF
			tailRet = tailRet.Next
			headF = tmp
			continue
		}

		if headF.Val > headS.Val {
			tmp := headS.Next
			headS.Next = nil
			tailRet.Next = headS
			tailRet = tailRet.Next
			headS = tmp
			continue
		}

		tmp := headF.Next
		headF.Next = nil
		tailRet.Next = headF
		tailRet = tailRet.Next
		headF = tmp

		tmp1 := headS.Next
		headS.Next = nil
		tailRet.Next = headS
		tailRet = tailRet.Next
		headS = tmp1
	}

	for headF != nil {
		tmp := headF.Next
		headF.Next = nil
		tailRet.Next = headF
		tailRet = tailRet.Next
		headF = tmp
	}

	for headS != nil {
		tmp := headS.Next
		headS.Next = nil
		tailRet.Next = headS
		tailRet = tailRet.Next
		headS = tmp
	}

	return headRet.Next
}

func main() {
	f := newList([]int{1, 10, 99, 1000})
	s := newList([]int{3, 6, 100, 123,345, 789})

	fmt.Println(displayList(f))
	fmt.Println(displayList(s))

	fmt.Println(displayList(mergeSortedLinkedList(f, s)))
}

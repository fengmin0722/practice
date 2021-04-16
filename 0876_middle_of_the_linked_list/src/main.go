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

	head := &ListNode{Val:nums[0], Next: nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		cur.Next = node
		cur = node
	}
	return head
}

func display(head *ListNode) string {
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

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	total := 0
	cur := head
	for cur != nil {
		total++
		cur = cur.Next
	}

	mid := total/2
	fmt.Println(fmt.Sprintf("total[%v], mid[%v]", total, mid))
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}

func middleNodeEx(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fmt.Println(fmt.Sprintf("s[%v], f[%v]", slow.Val, fast.Val))
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	list := newList([]int{1, 10, 100, 100})
	fmt.Println(display(list))
	fmt.Println(middleNode(list).Val)
	fmt.Println(display(list))
	fmt.Println(middleNodeEx(list).Val)
}

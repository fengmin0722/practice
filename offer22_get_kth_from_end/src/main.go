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
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		cur.Next = node
		cur = cur.Next
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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}
	total := 0
	cur := head
	for cur != nil {
		total++
		cur = cur.Next
	}
	cur = head
	sk := total-k
	for cur != nil && sk > 0 {
		sk--
		cur = cur.Next
	}
	return cur
}

func main() {
	list := newList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(displayList(list))
	fmt.Println(getKthFromEnd(list, 2).Val)
}

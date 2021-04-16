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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{head.Val, nil}
	cur := head.Next
	for cur != nil {
		newHead = prepend(newHead, cur.Val)
		cur = cur.Next
	}
	return newHead
}

func prepend(head *ListNode, Val int) *ListNode {
	node := &ListNode{Val: Val, Next: nil}
	if head == nil {
		return node
	}
	node.Next = head
	return node
}

func main() {
	list := newList([]int{1, 2, 3})
	fmt.Println(displayList(list))
	reList := reverseList(list)
	fmt.Println(displayList(reList))
}

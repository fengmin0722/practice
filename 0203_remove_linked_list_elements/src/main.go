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

func removeLinkedListElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	headPreNode := &ListNode{Next: head}
	pre, cur := headPreNode, head
	for cur != nil {
		if cur.Val == val {
			// 断开链路，移动当前指针
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre = pre.Next
		cur = cur.Next
	}
	return headPreNode.Next
}

func main() {
	list := newList([]int{1, 2, 3})
	fmt.Println(display(list))
	removeLinkedListElements(list, 2)
	fmt.Println(display(list))
}

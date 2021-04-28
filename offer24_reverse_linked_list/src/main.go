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
	head := &ListNode{Val:nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val:nums[i]}
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

func reverseDisplayList(head *ListNode) string {
	var res strings.Builder
	if head == nil {
		return res.String()
	}
	res.WriteString(reverseDisplayList(head.Next))
	res.WriteString(strconv.Itoa(head.Val))
	res.WriteString(" ")
	return res.String()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var dummyHead *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next // 为了遍历
		cur.Next = dummyHead
		dummyHead = cur
		cur = temp
	}
	return dummyHead
}

func main() {
	list := newList([]int{1, 2, 3, 4, 5})
	fmt.Println(displayList(list))
	fmt.Println(reverseDisplayList(list))

	rList := reverseList(list)
	fmt.Println(displayList(rList))
	fmt.Println(reverseDisplayList(rList))
}

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
	head := &ListNode{Val:nums[0], Next:nil}
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

func getPreAddr(dummyHead *ListNode, target *ListNode) *ListNode {
	cur := dummyHead
	for cur.Next != nil && cur.Next.Val < target.Val {
		cur = cur.Next
	}
	return cur
}

func insertSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//先把两个链表分开
	dummyHead := &ListNode{-1, head}
	cur := head.Next
	head.Next = nil
	//遍历链表
	for cur != nil {
		temp := cur.Next
		pre := getPreAddr(dummyHead, cur)
		cur.Next = pre.Next
		pre.Next = cur
		cur = temp
	}
	return dummyHead.Next
}

func main() {
	list := newList([]int{10, 8, 30, 2, 33})
	fmt.Println(displayList(list))
	fmt.Println(displayList(insertSort(list)))
	fmt.Println("ddd")
}

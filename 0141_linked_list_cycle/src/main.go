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

func hasCycle(head *ListNode) bool {
	// 空链表和只有一个节点的链表
	if head == nil || head.Next == nil {
		return false
	}
	addresses := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if _, ok := addresses[cur]; ok {
			return true
		}
		addresses[cur] = true
		cur = cur.Next
	}
	return false
}

func main() {
	list := newList([]int{1, 2, 3, 4})
	fmt.Println(displayList(list))
	fmt.Println(hasCycle(list))
	list.Next.Next.Next.Next = list.Next
	fmt.Println(hasCycle(list))
}

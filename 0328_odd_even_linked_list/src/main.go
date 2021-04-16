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
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val:nums[i]}
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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next // 偶数
	odd := head			  // 奇数
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	list := newList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(displayList(list))
	fmt.Println(displayList(oddEvenList(list)))
}

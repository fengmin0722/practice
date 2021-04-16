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
		node := &ListNode{Val:nums[i], Next: nil}
		cur.Next = node
		cur = node
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

func addTwoNums(left *ListNode, right *ListNode) *ListNode {
	if left == nil || right == nil {
		return nil
	}

	tmp := 0
	dummyHead := &ListNode{-1, nil}
	cur := dummyHead
	for left != nil || right != nil {
		var x int
		var y int
		if left != nil {
			x = left.Val
		}
		if right != nil {
			y = right.Val
		}

		total := x + y + tmp
		tmp = total / 10

		node := &ListNode{total % 10, nil}
		cur.Next = node
		cur = node

		if left != nil {
			left = left.Next
		}
		if right != nil {
			right = right.Next
		}
	}

	if tmp > 0 {
		cur.Next = &ListNode{tmp, nil}
	}

	return dummyHead.Next
}

func main() {
	listLeft := newList([]int{2, 4, 3})
	listRight := newList([]int{5, 6, 4})
	fmt.Println(displayList(listLeft))
	fmt.Println(displayList(listRight))
	fmt.Println(displayList(addTwoNums(listLeft, listRight)))
}

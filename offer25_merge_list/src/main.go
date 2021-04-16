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

func MergeList(left *ListNode, right *ListNode) *ListNode {
	if left == nil && right == nil {
		return nil
	}

	if left != nil && right == nil {
		return left
	}

	if left == nil && right != nil {
		return right
	}

	dummyNode := &ListNode{}
	cur := dummyNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			if dummyNode.Next == nil {
				dummyNode.Next = left
				cur = dummyNode.Next
			} else {
				cur.Next = left
				cur = cur.Next
			}
			left = left.Next
		} else {
			if dummyNode.Next == nil {
				dummyNode.Next = right
				cur = dummyNode.Next
			} else {
				cur.Next = right
				cur = cur.Next
			}
			right = right.Next
		}
	}

	if left != nil {
		cur.Next = left
	}

	if right != nil {
		cur.Next = right
	}

	return dummyNode.Next
}

func main() {
	listLeft := newList([]int{1, 3, 5})
	listRight := newList([]int{2, 4, 6})

	fmt.Println(displayList(listLeft))
	fmt.Println(displayList(listRight))

	fmt.Println(displayList(MergeList(listLeft, listRight)))
}

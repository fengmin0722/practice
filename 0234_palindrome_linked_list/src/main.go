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

func isPalindromeEx(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	total := 0
	cur := head
	for cur != nil {
		total++
		cur = cur.Next
	}

	mid := total / 2
	newHead := &ListNode{head.Val, nil}
	cur = head//
	for i := 0; i < mid - 1; i++ {
		cur = cur.Next
		node := &ListNode{cur.Val, newHead}
		newHead = node
	}

	cur = cur.Next
	if total % 2 == 1 {
		cur = cur.Next
	}

	for newHead != nil && cur != nil {
		if newHead.Val != cur.Val {
			return false
		}
		cur = cur.Next
		newHead = newHead.Next
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true // 注意 []、[1] 是回文结构
	}

	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	mid := n / 2 // 找到中间节点

	cur = head
	newHead := &ListNode{Val: head.Val} // 反转链表前半部分
	i := 0
	for i < mid-1 {
		cur = cur.Next
		newHead = prepend(newHead, cur.Val)
		i++
	}

	fmt.Println("cur ", cur.Val, "i ", i)
	cur = cur.Next // 选取要开始遍历的后半段链表
	if n%2 == 1 {
		cur = cur.Next
	}

	for newHead != nil && cur != nil {
		if newHead.Val != cur.Val {
			return false
		}
		cur = cur.Next
		newHead = newHead.Next
	}

	return true // 对称
}

// 新增 head
func prepend(head *ListNode, v int) (newHead *ListNode) {
	node := &ListNode{Val: v}
	if head == nil {
		return node
	}
	node.Next = head
	return node
}

func main() {
	list := newList([]int{2, 3, 3, 2})
	newList := newList([]int{1, 2, 1})
	fmt.Println(displayList(list))
	fmt.Println(isPalindromeEx(list))
	fmt.Println(isPalindromeEx(newList))

	fmt.Println(5/2)
}

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

func removeNthNode(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	front, rear := dummyNode, dummyNode
	for counts := 0; counts <= k; counts++ {
		rear = rear.Next
	}
	fmt.Println(fmt.Sprintf("rear [%v]", rear.Val))
	for rear != nil {
		front = front.Next
		rear = rear.Next
	}
	front.Next = front.Next.Next // 删除节点
	return dummyNode.Next
}

func main() {
	list := newList([]int{1, 2, 3, 4, 5})
	fmt.Println(displayList(list))
	fmt.Println(displayList(removeNthNode(list, 2)))
}

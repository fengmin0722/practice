package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func array2Tree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	root := &TreeNode{Val:nums[0]}
	queue := make([]*TreeNode, 1, len(nums))
	queue[0] = root

	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			node.Left = &TreeNode{Val:nums[i]}
			queue = append(queue, node.Left)
			i++
		}

		if i < len(nums) {
			node.Right = &TreeNode{Val:nums[i]}
			queue = append(queue, node.Right)
			i++
		}
	}
	return root
}

func findBottomLeftValue1(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) == 0 {
			return queue[0].Val
		}
		queue = next
	}
	return 0
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	root := array2Tree(nums)
	fmt.Println(findBottomLeftValue1(root))
}

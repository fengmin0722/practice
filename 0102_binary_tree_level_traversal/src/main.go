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

func levelOrder(root *TreeNode) []int {
	res := []int{}
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++{
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			res = append(res, node.Val)
		}
		queue = queue[l:]
	}

	return res
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	root := array2Tree(nums)
	fmt.Println(levelOrder(root))
}

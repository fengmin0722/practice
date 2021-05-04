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

		node.Left = &TreeNode{Val:nums[i]}
		queue = append(queue, node.Left)
		i++

		node.Right = &TreeNode{Val:nums[i]}
		queue = append(queue, node.Right)
		i++
	}
	return root
}

func inorder(root *TreeNode) []int {
	res := []int{}
	inorder_r(root, &res)
	return res
}

func inorder_r(root *TreeNode, output *[]int) {
	if root != nil {
		inorder_r(root.Left, output)
		*output = append(*output, root.Val)
		inorder_r(root.Right, output)
	}
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	root := array2Tree(nums)
	fmt.Println(inorder(root))
}

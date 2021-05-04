package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func array2Tree(nums []int) *TreeNode {
	if len(nums) == 0 {
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

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left)
		for _, v := range tmp {
			res = append(res, v)
		}
		rtmp := preorderTraversal(root.Right)
		for _, item := range rtmp {
			res = append(res, item)
		}
	}

	return res
}

func inorder(root *TreeNode, output  *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	root := array2Tree(nums)
	fmt.Println(preorderTraversal(root))
	fmt.Println(inorderTraversal(root))
}

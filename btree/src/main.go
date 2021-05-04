package main

import "fmt"

type Node struct {
	Data int
	Left *Node
	Right *Node
}

type BTree struct {
	Root *Node
}

func (this *BTree) find(data int) *Node {
	p := this.Root
	for p != nil {
		if p.Data == data {
			return p
		}
		if data < p.Data {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}

func (this *BTree) find_r(root *Node, data int) *Node {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	}
	if data < root.Data {
		return this.find_r(root.Left, data)
	} else {
		return this.find_r(root.Right, data)
	}
}

func (this *BTree) insert(data int) *Node {
	if this.Root == nil {
		this.Root = &Node{Data:data}
		return this.Root
	}

	if data > this.Root.Data {
		if this.Root.Right == nil {
			this.Root.Right = &Node{Data:data}
			return this.Root.Right
		} else {
			return this.insert_r(this.Root.Right, data)
		}
	} else {
		if this.Root.Left == nil {
			this.Root.Left = &Node{Data:data}
			return this.Root.Left
		} else {
			return this.insert_r(this.Root.Left, data)
		}
	}
}

func (this *BTree) insert_r(root *Node, data int) *Node {
	if root == nil {
		return nil
	}

	if data > root.Data {
		if root.Right == nil {
			root.Right = &Node{Data:data}
			return root.Right
		} else {
			return this.insert_r(root.Right, data)
		}
	} else {
		if root.Left == nil {
			root.Left = &Node{Data:data}
			return root.Left
		} else {
			return this.insert_r(root.Left, data)
		}
	}
}

func (this *BTree) insertNode(data int) {
	if this.Root == nil {
		this.Root = &Node{Data:data}
		return
	}

	p := this.Root
	for p != nil {
		if data > p.Data {
			if p.Right == nil {
				p.Right = &Node{Data:data}
				return
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = &Node{Data:data}
				return
			}
			p = p.Left
		}
	}
}

func (this *BTree) delete(data int) {
	p := this.Root
	var pp *Node
	for p != nil && p.Data != data {
		pp = p
		if data > p.Data {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	// 没有找到删除的节点
	if p == nil {
		return
	}

	// 删除有两个子节点
	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := minP
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Data = minP.Data
		p = minP
		pp = minPP
	}

	// 删除
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}

	if pp == nil {
		this.Root = child
	} else if pp.Left == p {
		pp.Left = child
	} else if pp.Right == p {
		pp.Right = child
	}
}

func main() {

	fmt.Println("hello")
}
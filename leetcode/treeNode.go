package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(nodes []int) (root *TreeNode) {
	root = &TreeNode{Val: nodes[0]}
	for i := range nodes {
		child := &TreeNode{Val: nodes[i]}
		child.InsertNodeToTree(root)
	}

	return
}

func (node *TreeNode) InsertNodeToTree(tree *TreeNode) {
	if tree == nil {
		return
	}

	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = node
		} else {
			node.InsertNodeToTree(tree.Left)
		}
	}

	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = node
		} else {
			node.InsertNodeToTree(tree.Right)
		}

	}
}

func (node *TreeNode) PrintTree() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	node.Left.PrintTree()
	node.Right.PrintTree()
}

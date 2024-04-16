package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func AddOneRow(root *utils.TreeNode, val int, depth int) *utils.TreeNode {
	if depth == 1 {
		return &utils.TreeNode{Val: val, Left: root}
	}

	var dfs func(node *utils.TreeNode, d int)
	dfs = func(node *utils.TreeNode, d int) {
		if node == nil {
			return
		}

		left, right := node.Left, node.Right
		if d-1 == 1 {
			node.Left = &utils.TreeNode{Val: val, Left: left}
			node.Right = &utils.TreeNode{Val: val, Right: right}
		}

		dfs(node.Left, d-1)
		dfs(node.Right, d-1)
	}

	dfs(root, depth)
	return root
}

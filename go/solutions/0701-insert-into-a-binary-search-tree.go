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
func InsertIntoBST(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return &utils.TreeNode{Val: val}
	}

	if root.Val > val {
		if root.Left == nil {
			root.Left = &utils.TreeNode{Val: val}
		} else {
			InsertIntoBST(root.Left, val)
		}
	}

	if root.Val < val {
		if root.Right == nil {
			root.Right = &utils.TreeNode{Val: val}
		} else {
			InsertIntoBST(root.Right, val)
		}
	}

	return root
}

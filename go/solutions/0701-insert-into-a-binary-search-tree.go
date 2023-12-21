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

	curr := root
	for {
		if curr.Val > val {
			if curr.Left == nil {
				curr.Left = &utils.TreeNode{Val: val}
				break
			} else {
				curr = curr.Left
			}
		} else {
			if curr.Right == nil {
				curr.Right = &utils.TreeNode{Val: val}
				break
			} else {
				curr = curr.Right
			}
		}
	}

	return root
}

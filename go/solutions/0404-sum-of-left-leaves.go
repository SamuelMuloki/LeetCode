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
func SumOfLeftLeaves(root *utils.TreeNode) int {
	return dfs2(root, false)
}

func dfs2(root *utils.TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return dfs2(root.Left, true) + dfs2(root.Right, false)
}

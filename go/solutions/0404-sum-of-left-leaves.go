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
	var dfs func(root *utils.TreeNode, isLeft bool) int
	dfs = func(root *utils.TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}

		if root.Left == nil && root.Right == nil && isLeft {
			return root.Val
		}

		return dfs(root.Left, true) + dfs(root.Right, false)
	}

	return dfs(root, false)
}

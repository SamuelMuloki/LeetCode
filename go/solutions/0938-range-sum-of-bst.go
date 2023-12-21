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
func RangeSumBST(root *utils.TreeNode, low int, high int) int {
	ans := 0
	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= low && node.Val <= high {
			ans += node.Val
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

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
func PseudoPalindromicPaths(root *utils.TreeNode) int {
	ans := 0
	var dfs func(node *utils.TreeNode, path int)
	dfs = func(node *utils.TreeNode, path int) {
		if node == nil {
			return
		}

		path = path ^ (1 << node.Val)
		if node.Left == nil && node.Right == nil {
			if path&(path-1) == 0 {
				ans++
			}
		}

		dfs(node.Left, path)
		dfs(node.Right, path)
	}

	dfs(root, 0)
	return ans
}

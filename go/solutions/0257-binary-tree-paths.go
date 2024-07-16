package solutions

import (
	"strconv"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BinaryTreePaths(root *utils.TreeNode) []string {
	var ans []string
	var dfs func(node *utils.TreeNode, prefix string)
	dfs = func(node *utils.TreeNode, prefix string) {
		if node == nil {
			return
		}

		if len(prefix) == 0 {
			prefix += strconv.Itoa(node.Val)
		} else {
			prefix += "->" + strconv.Itoa(node.Val)
		}

		if node.Left == nil && node.Right == nil {
			ans = append(ans, prefix)
		}

		dfs(node.Left, prefix)
		dfs(node.Right, prefix)
	}

	dfs(root, "")
	return ans
}

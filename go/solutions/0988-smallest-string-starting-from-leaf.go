package solutions

import (
	"strings"

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
func SmallestFromLeaf(root *utils.TreeNode) string {
	ans := ""
	var dfs func(node *utils.TreeNode, str string)
	dfs = func(node *utils.TreeNode, str string) {
		if node == nil {
			return
		}

		str = string(byte(node.Val)+'a') + str
		if node.Left == nil && node.Right == nil {
			if ans == "" || strings.Compare(str, ans) < 0 {
				ans = str
			}
		}

		dfs(node.Left, str)
		dfs(node.Right, str)
	}
	dfs(root, "")
	return ans
}

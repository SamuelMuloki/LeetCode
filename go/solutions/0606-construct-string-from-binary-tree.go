package solutions

import (
	"strconv"
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
func Tree2str(root *utils.TreeNode) string {
	ans := []string{}
	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		ans = append(ans, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			return
		}

		ans = append(ans, "(")
		dfs(node.Left)
		ans = append(ans, ")")

		if node.Right != nil {
			ans = append(ans, "(")
			dfs(node.Right)
			ans = append(ans, ")")
		}
	}

	dfs(root)
	return strings.Join(ans, "")
}

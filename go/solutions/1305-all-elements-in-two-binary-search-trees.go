package solutions

import (
	"sort"

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
func GetAllElements(root1 *utils.TreeNode, root2 *utils.TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}

	dfs(root1)
	dfs(root2)

	sort.Ints(ans)
	return ans
}

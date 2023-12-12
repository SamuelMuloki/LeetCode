package solutions

import (
	"math"

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
func MaxAncestorDiff(root *utils.TreeNode) int {
	ans := math.MinInt
	var dfs func(node *utils.TreeNode, minVal, maxVal int)
	dfs = func(node *utils.TreeNode, minVal, maxVal int) {
		if node == nil {
			return
		}

		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)

		dfs(node.Left, minVal, maxVal)
		dfs(node.Right, minVal, maxVal)

		ans = max(ans, maxVal-minVal)
	}

	dfs(root, math.MaxInt, math.MinInt)
	return ans
}

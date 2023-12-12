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
	var dfs func(node *utils.TreeNode, minVal, maxVal int) int
	dfs = func(node *utils.TreeNode, minVal, maxVal int) int {
		if node == nil {
			return 0
		}

		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)

		return max(maxVal-minVal, dfs(node.Left, minVal, maxVal), dfs(node.Right, minVal, maxVal))
	}

	return dfs(root, math.MaxInt, math.MinInt)
}

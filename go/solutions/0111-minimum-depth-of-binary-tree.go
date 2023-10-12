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
func MinDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	minVal := math.MaxInt
	var dfs func(root *utils.TreeNode, count int)
	dfs = func(root *utils.TreeNode, count int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			minVal = utils.Min(minVal, count)
			return
		}

		count += 1
		dfs(root.Left, count)
		dfs(root.Right, count)
	}

	dfs(root, 1)

	return minVal
}

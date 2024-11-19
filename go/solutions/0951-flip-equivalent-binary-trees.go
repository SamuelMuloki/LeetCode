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
func FlipEquiv(root1 *utils.TreeNode, root2 *utils.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	noSwap := FlipEquiv(root1.Left, root2.Left) &&
		FlipEquiv(root1.Right, root2.Right)
	swap := FlipEquiv(root1.Left, root2.Right) &&
		FlipEquiv(root1.Right, root2.Left)

	return swap || noSwap
}

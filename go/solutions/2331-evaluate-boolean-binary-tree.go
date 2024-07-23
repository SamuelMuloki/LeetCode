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
func EvaluateTree(root *utils.TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val != 0
	}

	evaluateLeftSubtree := EvaluateTree(root.Left)
	evaluateRightSubtree := EvaluateTree(root.Right)
	evaluateRoot := false
	if root.Val == 2 {
		evaluateRoot = evaluateLeftSubtree || evaluateRightSubtree
	} else {
		evaluateRoot = evaluateLeftSubtree && evaluateRightSubtree
	}

	return evaluateRoot
}

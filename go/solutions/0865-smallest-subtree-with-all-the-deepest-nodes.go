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
func SubtreeWithAllDeepest(root *utils.TreeNode) *utils.TreeNode {
	deepest := 0
	ans := new(utils.TreeNode)
	var helper func(node *utils.TreeNode, depth int) int
	helper = func(node *utils.TreeNode, depth int) int {
		deepest = max(deepest, depth)
		if node == nil {
			return depth
		}

		left := helper(node.Left, depth+1)
		right := helper(node.Right, depth+1)
		if left == deepest && right == deepest {
			ans = node
		}

		return max(left, right)
	}

	helper(root, 0)
	return ans
}

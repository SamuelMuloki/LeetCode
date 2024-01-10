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
func AmountOfTime(root *utils.TreeNode, start int) int {
	ans := 0
	var dfs func(node *utils.TreeNode) int
	dfs = func(node *utils.TreeNode) int {
		depth := 0
		if node == nil {
			return depth
		}

		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)

		if node.Val == start {
			ans = max(leftDepth, rightDepth)
			depth = -1
		} else if leftDepth >= 0 && rightDepth >= 0 {
			depth = max(leftDepth, rightDepth) + 1
		} else {
			distance := abs(leftDepth) + abs(rightDepth)
			ans = max(ans, distance)
			depth = min(leftDepth, rightDepth) - 1
		}

		return depth
	}

	dfs(root)
	return ans
}

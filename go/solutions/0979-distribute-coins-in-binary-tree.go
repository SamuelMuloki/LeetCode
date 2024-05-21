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
func DistributeCoins(root *utils.TreeNode) int {
	var moves int

	var dfs func(node *utils.TreeNode) int
	dfs = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}
		leftExcess := dfs(node.Left)
		rightExcess := dfs(node.Right)
		totalExcess := node.Val + leftExcess + rightExcess - 1
		moves += abs(totalExcess)
		return totalExcess
	}

	dfs(root)
	return moves
}

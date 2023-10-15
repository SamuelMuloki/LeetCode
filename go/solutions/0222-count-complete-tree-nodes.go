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
func CountNodes(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	res := 1
	if leftHeight == rightHeight {
		res += (1 << leftHeight) - 1
		res += CountNodes(root.Right)
	} else {
		res += (1 << rightHeight) - 1
		res += CountNodes(root.Left)
	}
	return res
}

func getHeight(node *utils.TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + getHeight(node.Left)
}

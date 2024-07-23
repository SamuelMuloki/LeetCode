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
func RemoveLeafNodes(root *utils.TreeNode, target int) *utils.TreeNode {
	if root == nil {
		return nil
	}

	root.Left = RemoveLeafNodes(root.Left, target)
	root.Right = RemoveLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

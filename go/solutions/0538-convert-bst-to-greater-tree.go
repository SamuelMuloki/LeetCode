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
func ConvertBST(root *utils.TreeNode) *utils.TreeNode {
	val := 0
	return recur(root, &val)
}

func recur(root *utils.TreeNode, val *int) *utils.TreeNode {
	if root == nil {
		return root
	}

	recur(root.Right, val)
	*val += root.Val
	root.Val = *val
	recur(root.Left, val)

	return root
}

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
func CheckTree(root *utils.TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

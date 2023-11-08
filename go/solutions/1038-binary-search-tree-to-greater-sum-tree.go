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
func BstToGst(root *utils.TreeNode) *utils.TreeNode {
	val := 0
	return recur(root, &val)
}

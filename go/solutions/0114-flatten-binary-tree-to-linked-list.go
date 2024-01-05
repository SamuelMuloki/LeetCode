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
func Flatten(root *utils.TreeNode) {
	var prev *utils.TreeNode
	var flat func(root *utils.TreeNode)
	flat = func(root *utils.TreeNode) {
		if root == nil {
			return
		}

		flat(root.Right)
		flat(root.Left)
		root.Right = prev
		root.Left = nil

		prev = root
	}

	flat(root)
}

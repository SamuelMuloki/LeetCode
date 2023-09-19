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
func PreorderTraversal(root *utils.TreeNode) []int {
	return recur(root, &[]int{})
}

func recur(root *utils.TreeNode, output *[]int) []int {
	if root == nil {
		return *output
	}

	*output = append(*output, root.Val)
	recur(root.Left, output)
	recur(root.Right, output)

	return *output
}

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
func KthSmallest(root *utils.TreeNode, k int) int {
	if root == nil {
		return 0
	}

	output := make([]int, 0)
	recur(root, &output)

	return output[k-1]
}

func recur(root *utils.TreeNode, output *[]int) []int {
	if root == nil {
		return *output
	}

	recur(root.Left, output)
	*output = append(*output, root.Val)
	recur(root.Right, output)

	return *output
}

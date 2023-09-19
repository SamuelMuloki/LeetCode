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
func InorderTraversal(root *utils.TreeNode) []int {
	output := make([]int, 0)

	if root == nil {
		return output
	}

	s := make([]*utils.TreeNode, 0)
	curr := root

	for curr != nil || len(s) != 0 {
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}

		s, curr = s[:len(s)-1], s[len(s)-1]
		output = append(output, curr.Val)
		curr = curr.Right
	}

	return output
}

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
	output := make([]int, 0)

	if root == nil {
		return output
	}

	s := make([]*utils.TreeNode, 0)
	s = append(s, root)

	for len(s) != 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]

		output = append(output, last.Val)
		if last.Right != nil {
			s = append(s, last.Right)
		}

		if last.Left != nil {
			s = append(s, last.Left)
		}
	}

	return output
}

package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PostorderTraversal(root *utils.TreeNode) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}

	curr := root
	s := make([]*utils.TreeNode, 0)
	for curr != nil || len(s) != 0 {
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}

		curr = s[len(s)-1]
		curr = curr.Right

		// don't go down right branch if we've already been there
		if curr == nil || len(output) > 0 && curr.Val == output[len(output)-1] {
			s, curr = s[:len(s)-1], s[len(s)-1]
			output = append(output, curr.Val)
			curr = nil
		}
	}

	return output
}

package utils

import "github.com/SamuelMuloki/GOExamples/leetcode"

func CountTreeNodes(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	} else {
		return CountTreeNodes(root.Left) + CountTreeNodes(root.Right) + 1
	}
}

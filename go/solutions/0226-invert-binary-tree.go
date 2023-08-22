package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func InvertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

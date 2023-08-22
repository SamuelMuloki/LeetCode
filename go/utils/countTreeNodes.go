package utils

func CountTreeNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return CountTreeNodes(root.Left) + CountTreeNodes(root.Right) + 1
	}
}

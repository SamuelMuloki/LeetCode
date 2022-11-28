package leetcode

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsSymm(root.Left, root.Right)
}

func IsSymm(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil || y == nil || x.Val != y.Val {
		return false
	}

	return IsSymm(x.Left, y.Right) && IsSymm(x.Right, y.Left)
}

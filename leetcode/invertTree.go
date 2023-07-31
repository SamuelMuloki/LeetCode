package leetcode

func InvertTree(root *TreeNode) *TreeNode {
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

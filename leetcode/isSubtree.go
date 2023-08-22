package leetcode

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isIdentical(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isIdentical(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == nil && node2 == nil
	}

	return node1.Val == node2.Val && isIdentical(node1.Left, node2.Left) && isIdentical(node1.Right, node2.Right)
}

package leetcode

func MaxDepth(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}

	leftMaxDepth := MaxDepth(root.Left)
	rightMaxDepth := MaxDepth(root.Right)

	if leftMaxDepth > rightMaxDepth {
		count = leftMaxDepth
	} else {
		count = rightMaxDepth
	}

	return count + 1
}

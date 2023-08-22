package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaxDepth(root *utils.TreeNode) int {
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

package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func SumNumbers(root *utils.TreeNode) int {
	var dfs func(root *utils.TreeNode, sum int) int
	dfs = func(root *utils.TreeNode, sum int) int {
		if root == nil {
			return 0
		}

		val := sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return val
		}

		return dfs(root.Left, val) + dfs(root.Right, val)
	}

	return dfs(root, 0)
}

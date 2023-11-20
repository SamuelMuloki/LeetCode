package solutions

import (
	"strconv"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func SumNumbers(root *utils.TreeNode) int {
	sum := 0

	var dfs func(root *utils.TreeNode, sum *int, str string)
	dfs = func(root *utils.TreeNode, sum *int, str string) {
		if root == nil {
			return
		}

		str += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			num, _ := strconv.Atoi(str)
			*sum += num
			return
		}

		dfs(root.Left, sum, str)
		dfs(root.Right, sum, str)
	}

	dfs(root, &sum, "")
	return sum
}

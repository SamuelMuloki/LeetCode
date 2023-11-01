package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func FindMode(root *utils.TreeNode) []int {
	currStreak, maxStreak, currNum := 0, 0, 0
	ans := make([]int, 0)

	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if node.Val == currNum {
			currStreak++
		} else {
			currStreak = 1
			currNum = node.Val
		}

		if currStreak > maxStreak {
			ans = []int{}
			maxStreak = currStreak
		}

		if currStreak == maxStreak {
			ans = append(ans, node.Val)
		}

		dfs(node.Right)
	}

	dfs(root)

	return ans
}

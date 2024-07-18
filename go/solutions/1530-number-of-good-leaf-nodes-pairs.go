package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CountPairs3(root *utils.TreeNode, distance int) int {
	var dfs func(node *utils.TreeNode) []int
	dfs = func(node *utils.TreeNode) []int {
		if node == nil {
			return make([]int, 12)
		} else if node.Left == nil && node.Right == nil {
			current := make([]int, 12)
			current[0] = 1
			return current
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		current := make([]int, 12)
		for i := 0; i < 10; i++ {
			current[i+1] += left[i] + right[i]
		}

		current[11] += left[11] + right[11]
		for d1 := 0; d1 <= distance; d1++ {
			for d2 := 0; d2 <= distance; d2++ {
				if 2+d1+d2 <= distance {
					current[11] += left[d1] * right[d2]
				}
			}
		}

		return current
	}

	return dfs(root)[11]
}

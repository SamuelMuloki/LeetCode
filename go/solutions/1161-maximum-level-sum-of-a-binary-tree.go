package solutions

import (
	"math"

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
func MaxLevelSum(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)

	maxSum, maxLevel, level := math.MinInt, 0, 1
	for len(queue) > 0 {
		last, currSum := len(queue), 0
		for i := 0; i < last; i++ {
			currSum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		if currSum > maxSum {
			maxLevel = level
			maxSum = currSum
		}

		queue = queue[last:]
		level++
	}

	return maxLevel
}

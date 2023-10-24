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
func LargestValues(root *utils.TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		qLen, maxVal := len(queue), math.MinInt
		for i := 0; i < qLen; i++ {
			maxVal = utils.Max(maxVal, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		if maxVal != math.MinInt {
			ans = append(ans, maxVal)
		}
		queue = queue[qLen:]
	}

	return ans
}

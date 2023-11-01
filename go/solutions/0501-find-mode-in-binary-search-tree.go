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
func FindMode(root *utils.TreeNode) []int {
	ans := make([]int, 0)
	set := make(map[int]int)
	maxFreq := math.MinInt

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			set[queue[i].Val]++
			maxFreq = utils.Max(maxFreq, set[queue[i].Val])
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[qLen:]
	}

	for k := range set {
		if set[k] == maxFreq {
			ans = append(ans, k)
		}
	}

	return ans
}

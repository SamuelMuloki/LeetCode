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
func MinDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)

	level := 1
	for len(queue) > 0 {
		currLen := len(queue)
		for i := 0; i < currLen; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return level
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[currLen:]
		level++
	}

	return -1
}

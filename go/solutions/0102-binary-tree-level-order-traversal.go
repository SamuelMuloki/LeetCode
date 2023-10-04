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
type Level struct {
	level int
	node  *utils.TreeNode
}

func LevelOrder(root *utils.TreeNode) [][]int {
	output := make([][]int, 0)

	if root == nil {
		return output
	}

	queue := make([]Level, 0)
	queue = append(queue, Level{level: 0, node: root})

	for len(queue) > 0 {
		last := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if len(output)-1 >= last.level {
			output[last.level] = append(output[last.level], last.node.Val)
		} else {
			output = append(output, []int{last.node.Val})
		}

		if last.node.Right != nil {
			queue = append(queue, Level{level: last.level + 1, node: last.node.Right})
		}

		if last.node.Left != nil {
			queue = append(queue, Level{level: last.level + 1, node: last.node.Left})
		}
	}

	return output
}

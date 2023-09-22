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

type CurrNode struct {
	minVal, maxVal int
	node           *utils.TreeNode
}

func IsValidBST(root *utils.TreeNode) bool {
	queue := make([]CurrNode, 0)
	queue = append(queue, CurrNode{minVal: math.MinInt64, maxVal: math.MaxInt64, node: root})

	for len(queue) > 0 {
		last := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if last.node.Val <= last.minVal || last.node.Val >= last.maxVal {
			return false
		}

		if last.node.Right != nil {
			queue = append(queue, CurrNode{minVal: last.node.Val, maxVal: last.maxVal, node: last.node.Right})
		}

		if last.node.Left != nil {
			queue = append(queue, CurrNode{minVal: last.minVal, maxVal: last.node.Val, node: last.node.Left})
		}
	}

	return true
}

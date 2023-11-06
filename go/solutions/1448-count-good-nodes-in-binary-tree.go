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
func GoodNodes(root *utils.TreeNode) int {
	var count int
	st := make([]CurrNode, 0)
	st = append(st, CurrNode{minVal: math.MinInt64, node: root})
	for len(st) > 0 {
		last := st[len(st)-1]
		st = st[:len(st)-1]

		minNodeVal := last.minVal
		if last.node.Val >= last.minVal {
			minNodeVal = last.node.Val
			count++
		}

		if last.node.Right != nil {
			st = append(st, CurrNode{minVal: minNodeVal, node: last.node.Right})
		}

		if last.node.Left != nil {
			st = append(st, CurrNode{minVal: minNodeVal, node: last.node.Left})
		}
	}

	return count
}

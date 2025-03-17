package solutions

import (
	"strconv"

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
func RecoverFromPreorder(traversal string) *utils.TreeNode {
	levels := make([]*utils.TreeNode, 0)
	n := len(traversal)
	for i := 0; i < n; {
		depth := 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}

		numStr := ""
		for i < n && traversal[i] != '-' {
			numStr += strconv.Itoa(int(traversal[i] - '0'))
			i++
		}

		num, _ := strconv.Atoi(numStr)

		node := &utils.TreeNode{Val: num}
		if depth < len(levels) {
			levels[depth] = node
		} else {
			levels = append(levels, node)
		}

		if depth > 0 {
			parent := levels[depth-1]
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}
	}

	return levels[0]
}

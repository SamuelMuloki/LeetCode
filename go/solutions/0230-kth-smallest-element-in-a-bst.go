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
func KthSmallest(root *utils.TreeNode, k int) int {
	st := make([]*utils.TreeNode, 0)
	for root != nil {
		st = append(st, root)
		root = root.Left
	}

	for k != 0 {
		last := st[len(st)-1]
		st = st[:len(st)-1]
		k--
		if k == 0 {
			return last.Val
		}

		right := last.Right
		for right != nil {
			st = append(st, right)
			right = right.Left
		}
	}

	return -1
}

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
func SearchBST(root *utils.TreeNode, val int) *utils.TreeNode {
	st := make([]*utils.TreeNode, 0)
	st = append(st, root)

	for len(st) > 0 {
		last := st[len(st)-1]
		st = st[:len(st)-1]

		if last.Val == val {
			return last
		}

		if last.Right != nil {
			st = append(st, last.Right)
		}

		if last.Left != nil {
			st = append(st, last.Left)
		}
	}

	return nil
}

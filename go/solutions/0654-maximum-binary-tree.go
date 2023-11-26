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
func ConstructMaximumBinaryTree(nums []int) *utils.TreeNode {
	st := make([]*utils.TreeNode, 0)
	for i := range nums {
		node := &utils.TreeNode{Val: nums[i]}
		for len(st) > 0 && st[len(st)-1].Val < node.Val {
			node.Left = st[len(st)-1]
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			st[len(st)-1].Right = node
		}

		st = append(st, node)
	}

	for len(st) > 1 {
		st = st[:len(st)-1]
	}

	return st[len(st)-1]
}

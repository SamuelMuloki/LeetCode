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

type PathSum struct {
	currSum int
	node    *utils.TreeNode
}

func HasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	s := make([]PathSum, 0)
	s = append(s, PathSum{node: root, currSum: targetSum})

	for len(s) > 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]

		if last.node.Left == nil && last.node.Right == nil && last.currSum == last.node.Val {
			return true
		}

		if last.node.Right != nil {
			s = append(s, PathSum{node: last.node.Right, currSum: last.currSum - last.node.Val})
		}

		if last.node.Left != nil {
			s = append(s, PathSum{node: last.node.Left, currSum: last.currSum - last.node.Val})
		}
	}

	return false
}

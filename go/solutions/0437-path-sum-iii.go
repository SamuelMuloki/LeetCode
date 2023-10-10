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

func PathSum3(root *utils.TreeNode, targetSum int) int {
	var paths int
	s := make([]PathSum, 0)
	s = append(s, PathSum{node: root, currSum: targetSum})

	for len(s) > 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]

		st := make([]PathSum, 0)
		st = append(st, last)
		for len(st) > 0 {
			nodeLast := st[len(st)-1]
			st = st[:len(st)-1]

			if nodeLast.currSum == nodeLast.node.Val {
				paths++
			}

			if nodeLast.node.Right != nil {
				st = append(st, PathSum{node: nodeLast.node.Right, currSum: nodeLast.currSum - nodeLast.node.Val})
			}

			if nodeLast.node.Left != nil {
				st = append(st, PathSum{node: nodeLast.node.Left, currSum: nodeLast.currSum - nodeLast.node.Val})
			}
		}

		if last.node.Right != nil {
			s = append(s, PathSum{node: last.node.Right, currSum: targetSum})
		}

		if last.node.Left != nil {
			s = append(s, PathSum{node: last.node.Left, currSum: targetSum})
		}
	}

	return paths
}

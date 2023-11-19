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
func ReverseOddLevels(root *utils.TreeNode) *utils.TreeNode {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		if level%2 != 0 {
			for i, j := 0, qLen-1; i < j; i, j = i+1, j-1 {
				queue[i].Val, queue[j].Val = queue[j].Val, queue[i].Val
			}
		}

		queue = queue[qLen:]
		level++
	}

	return root
}

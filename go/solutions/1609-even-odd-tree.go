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
func IsEvenOddTree(root *utils.TreeNode) bool {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		qLen := len(queue)
		var prev int
		for i := 0; i < qLen; i++ {
			if level%2 == 0 {
				if queue[i].Val%2 != 1 || (i > 0 && prev >= queue[i].Val) {
					return false
				}
			} else {
				if queue[i].Val%2 != 0 || (i > 0 && prev <= queue[i].Val) {
					return false
				}
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			prev = queue[i].Val
		}

		queue = queue[qLen:]
		level++
	}

	return true
}

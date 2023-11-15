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
func ZigzagLevelOrder(root *utils.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue, level := make([]*utils.TreeNode, 0), 0
	queue = append(queue, root)
	for len(queue) > 0 {
		qLen, curr := len(queue), make([]int, 0)
		for i := 0; i < qLen; i++ {
			if level%2 == 0 {
				curr = append(curr, queue[i].Val)
			} else {
				curr = append([]int{queue[i].Val}, curr...)
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		ans = append(ans, append([]int{}, curr...))
		queue = queue[qLen:]
		level++
	}

	return ans
}

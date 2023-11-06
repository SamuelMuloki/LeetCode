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
func RightSideView(root *utils.TreeNode) []int {
	rightView := make([]int, 0)
	if root == nil {
		return rightView
	}

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		currLen := len(queue)
		rightView = append(rightView, queue[currLen-1].Val)
		for i := 0; i < currLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[currLen:]
	}

	return rightView
}

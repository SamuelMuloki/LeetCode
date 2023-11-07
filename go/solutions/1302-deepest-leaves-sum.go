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
func DeepestLeavesSum(root *utils.TreeNode) int {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		qLen := len(queue)
		ans = 0
		for i := 0; i < qLen; i++ {
			ans += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[qLen:]
	}

	return ans
}

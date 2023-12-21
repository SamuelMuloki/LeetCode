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
func FindBottomLeftValue(root *utils.TreeNode) int {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	ans := []int{}
	for len(queue) > 0 {
		qLen := len(queue)
		ans = []int{}
		for i := 0; i < qLen; i++ {
			ans = append(ans, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[qLen:]
	}

	return ans[0]
}

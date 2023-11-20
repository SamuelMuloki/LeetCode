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
func LevelOrderBottom(root *utils.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		first := len(queue)
		curr := make([]int, 0)
		for i := 0; i < first; i++ {
			curr = append(curr, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		ans = append([][]int{curr}, ans...)
		queue = queue[first:]
	}

	return ans
}

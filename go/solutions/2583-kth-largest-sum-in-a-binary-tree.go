package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func KthLargestLevelSum(root *utils.TreeNode, k int) int64 {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	arr := []int{}
	for len(queue) > 0 {
		qLen := len(queue)
		sum := 0
		for i := 0; i < qLen; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		arr = append(arr, sum)
		queue = queue[qLen:]
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
	if k-1 >= len(arr) {
		return -1
	}

	return int64(arr[k-1])
}

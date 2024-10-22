package solutions

import (
	"container/heap"

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
	h := &MinHeap{}
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

		heap.Push(h, sum)
		if h.Len() > k {
			heap.Pop(h)
		}
		queue = queue[qLen:]
	}

	if k > h.Len() {
		return -1
	}

	return int64((*h)[0])
}

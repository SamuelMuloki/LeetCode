package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func NodesBetweenCriticalPoints(head *utils.ListNode) []int {
	var prev int
	indices := []int{}
	for curr, idx := head, 0; curr != nil; curr = curr.Next {
		if curr.Next != nil && prev != 0 {
			if curr.Val < prev && curr.Val < curr.Next.Val ||
				curr.Val > prev && curr.Val > curr.Next.Val {
				indices = append(indices, idx)
			}
		}

		prev = curr.Val
		idx++
	}

	if len(indices) < 2 {
		return []int{-1, -1}
	}

	minDistance, n := math.MaxInt, len(indices)
	for i := 1; i < n; i++ {
		minDistance = min(minDistance, indices[i]-indices[i-1])
	}

	return []int{minDistance, indices[n-1] - indices[0]}
}

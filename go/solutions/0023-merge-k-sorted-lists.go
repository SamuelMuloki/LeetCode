package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*utils.ListNode) *utils.ListNode {
	var result []*utils.ListNode

	for i := 0; i < len(lists); i++ {
		var curr = lists[i]

		for curr != nil {
			result = append(result, curr)
			curr = curr.Next
		}
	}

	if len(result) == 0 {
		return nil
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Val < result[j].Val
	})

	for i := 1; i < len(result); i += 1 {
		result[i-1].Next = result[i]
	}

	if len(result) != 0 {
		result[len(result)-1].Next = nil
	}

	return result[0]
}

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
func SortList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	var ans []*utils.ListNode
	curr := head
	for curr != nil {
		ans = append(ans, curr)
		curr = curr.Next
	}

	if len(ans) == 0 {
		return nil
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Val < ans[j].Val
	})

	for i := 1; i < len(ans); i++ {
		ans[i-1].Next = ans[i]
	}

	if len(ans) != 0 {
		ans[len(ans)-1].Next = nil
	}

	return ans[0]
}

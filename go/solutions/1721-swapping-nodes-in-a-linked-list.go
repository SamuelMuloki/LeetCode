package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapNodes(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return nil
	}

	count, curr := 0, head
	for curr != nil {
		count++
		curr = curr.Next
	}

	begin := head
	for i := 1; i < k; i++ {
		begin = begin.Next
	}

	end := head
	for i := 0; i < count-k; i++ {
		end = end.Next
	}

	begin.Val, end.Val = end.Val, begin.Val

	return head
}

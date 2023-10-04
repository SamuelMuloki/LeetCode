package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(utils.ListNode)
	prev := dummy
	curr := head

	for curr != nil && curr.Next != nil {
		prev.Next = curr.Next
		curr.Next = prev.Next.Next
		prev.Next.Next = curr

		prev = curr
		curr = curr.Next
	}

	return dummy.Next
}

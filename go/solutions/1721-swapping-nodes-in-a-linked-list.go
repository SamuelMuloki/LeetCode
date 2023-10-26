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
	begin, end := head, head
	for curr := head; curr != nil; curr = curr.Next {
		k--
		if k == 0 {
			begin = curr
		}

		if k < 0 {
			end = end.Next
		}
	}

	begin.Val, end.Val = end.Val, begin.Val

	return head
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveElements(head *utils.ListNode, val int) *utils.ListNode {
	dummy := &utils.ListNode{Val: 0, Next: head}
	curr, prev := head, dummy
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		curr = curr.Next
	}

	return dummy.Next
}

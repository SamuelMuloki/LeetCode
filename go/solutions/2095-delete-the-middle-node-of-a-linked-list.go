package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteMiddle(head *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{Val: 0, Next: head}
	prev := head
	curr, count := head, 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	if count == 1 {
		return head.Next
	}

	mid := (count / 2)
	for i := 0; i < mid; i++ {
		if i+1 == mid {
			prev.Next = prev.Next.Next
		}

		prev = prev.Next
	}

	return dummy.Next
}

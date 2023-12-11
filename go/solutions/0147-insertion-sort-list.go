package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func InsertionSortList(head *utils.ListNode) *utils.ListNode {
	dummy := new(utils.ListNode)
	curr := head
	for curr != nil {
		prev := dummy
		for prev.Next != nil && prev.Next.Val <= curr.Val {
			prev = prev.Next
		}

		next := curr.Next
		curr.Next = prev.Next
		prev.Next = curr

		curr = next
	}

	return dummy.Next
}

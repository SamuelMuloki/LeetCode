package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func PartitionList(head *utils.ListNode, x int) *utils.ListNode {
	ans := &utils.ListNode{Val: 0}

	curr, dummy := head, ans
	for curr != nil {
		if curr.Val < x {
			dummy.Next = &utils.ListNode{Val: curr.Val}
			dummy = dummy.Next
		}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		if curr.Val >= x {
			dummy.Next = &utils.ListNode{Val: curr.Val}
			dummy = dummy.Next
		}
		curr = curr.Next
	}

	return ans.Next
}

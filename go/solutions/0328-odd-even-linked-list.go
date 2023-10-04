package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func OddEvenList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	dummy := &utils.ListNode{Val: 0, Next: head}
	slow, fast := head, head.Next
	fastDummy := fast
	for fast != nil && fast.Next != nil {
		slow.Next = slow.Next.Next
		fast.Next = fast.Next.Next

		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = fastDummy

	return dummy.Next
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveZeroSumSublists(head *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{Val: 0, Next: head}

	m := make(map[int]*utils.ListNode)
	m[0] = dummy

	cur := head
	sum := 0
	for cur != nil {
		sum += cur.Val
		m[sum] = cur
		cur = cur.Next
	}

	sum = 0
	cur = dummy
	for cur != nil {
		sum += cur.Val
		cur.Next = m[sum].Next
		cur = cur.Next
	}

	return dummy.Next
}

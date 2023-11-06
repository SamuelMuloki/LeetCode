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
	l1, l2 := new(utils.ListNode), new(utils.ListNode)
	d1, d2 := l1, l2
	for head != nil {
		if head.Val < x {
			l1.Next = head
			l1 = l1.Next
		} else {
			l2.Next = head
			l2 = l2.Next
		}
		head = head.Next
	}

	l2.Next = nil
	l1.Next = d2.Next

	return d1.Next
}

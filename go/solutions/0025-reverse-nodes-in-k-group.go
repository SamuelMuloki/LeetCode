package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	curr, count := head, 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	ptr, dummy := head, new(utils.ListNode)
	d2 := dummy
	for i := 0; i < count/k; i++ {
		for j := 0; j < k; j++ {
			tmp := ptr.Next
			ptr.Next = d2.Next
			d2.Next = ptr
			ptr = tmp
		}

		for d2.Next != nil {
			d2 = d2.Next
		}
	}
	d2.Next = ptr

	return dummy.Next
}

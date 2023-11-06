package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr, count := head, 0
	for ; curr != nil; curr = curr.Next {
		count++
	}

	end, start, last := head, head, k%count
	for i := 0; i < count-last; i++ {
		end = end.Next
	}

	for i := 1; i < count-last; i++ {
		start = start.Next
	}

	start.Next = nil
	newEnd := end
	for ; newEnd != nil && newEnd.Next != nil; newEnd = newEnd.Next {
	}

	if newEnd != nil {
		newEnd.Next = head
	} else {
		return head
	}

	return end
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *utils.ListNode
 * }
 */
func ReverseBetween(head *utils.ListNode, left int, right int) *utils.ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &utils.ListNode{Val: 0, Next: head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next

	for i := 0; i < right-left; i++ {
		nextNode := current.Next
		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}

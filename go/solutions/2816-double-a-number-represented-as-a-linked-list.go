package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DoubleIt(head *utils.ListNode) *utils.ListNode {
	head = &utils.ListNode{Next: head}
	for curr, next := head, head.Next; next != nil; curr, next = next, next.Next {
		next.Val *= 2
		curr.Val += next.Val / 10
		next.Val %= 10
	}

	if head.Val == 0 {
		head = head.Next
	}

	return head
}

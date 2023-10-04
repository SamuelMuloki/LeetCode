package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func ReverseList(head *utils.ListNode) *utils.ListNode {
	curr := head
	var prev *utils.ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

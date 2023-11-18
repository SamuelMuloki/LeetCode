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
func PairSum(head *utils.ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second, prev := slow, new(utils.ListNode)
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	maxSum := 0
	for head != slow && prev != nil {
		maxSum = max(maxSum, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}

	return maxSum
}

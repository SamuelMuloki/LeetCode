package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	end, l, curr := head, 0, head
	for curr != nil {
		l++
		curr = curr.Next
	}

	prev := end
	if n == l {
		prev = prev.Next
		return prev
	}

	for i := 0; i < l; i++ {
		if i == l-n {
			prev.Next = end.Next
			break
		}

		prev = end
		end = end.Next
	}

	return head
}

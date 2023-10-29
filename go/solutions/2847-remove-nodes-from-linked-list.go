package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNodes(head *utils.ListNode) *utils.ListNode {
	st := make([]*utils.ListNode, 0)
	curr := head
	for curr != nil {
		for len(st) > 0 && st[len(st)-1].Val < curr.Val {
			st = st[:len(st)-1]
		}

		st = append(st, curr)
		curr = curr.Next
	}

	ans := new(utils.ListNode)
	dummy := ans
	for i := 0; i < len(st); i++ {
		dummy.Next = st[i]
		dummy = dummy.Next
	}

	return ans.Next
}

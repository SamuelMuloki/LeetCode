package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeInBetween(list1 *utils.ListNode, a int, b int, list2 *utils.ListNode) *utils.ListNode {
	curr, prev := list1, new(utils.ListNode)
	count := 0
	for ; curr != nil; curr = curr.Next {
		if count == a {
			prev.Next = list2
		}

		if count == b {
			break
		}

		prev = curr
		count++
	}

	for ; list2.Next != nil; list2 = list2.Next {
	}
	list2.Next = curr.Next

	return list1
}

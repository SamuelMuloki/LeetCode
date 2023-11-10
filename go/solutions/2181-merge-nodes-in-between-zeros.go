package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeNodes(head *utils.ListNode) *utils.ListNode {
	currSum, dummy := 0, new(utils.ListNode)
	d2 := dummy
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == 0 && currSum > 0 {
			d2.Next = &utils.ListNode{Val: currSum}
			d2, currSum = d2.Next, 0
		} else {
			currSum += curr.Val
		}
	}

	return dummy.Next
}

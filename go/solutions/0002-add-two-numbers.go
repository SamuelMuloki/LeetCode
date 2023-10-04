package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	dummyPtr, carry := new(utils.ListNode), 0
	final := dummyPtr

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		dummyPtr.Next = &utils.ListNode{Val: sum % 10}
		dummyPtr = dummyPtr.Next
	}

	return final.Next
}

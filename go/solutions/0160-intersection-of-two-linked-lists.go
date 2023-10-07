package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	currA, currB := headA, headB
	for currA != nil && currB != nil {
		tempB := currB
		for tempB != nil && currA.Val != tempB.Val {
			tempB = tempB.Next
		}

		if currA == tempB {
			return currA
		}

		currA = currA.Next
	}

	return nil
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ModifiedList(nums []int, head *utils.ListNode) *utils.ListNode {
	delSet := make(map[int]int)
	for i := range nums {
		delSet[nums[i]]++
	}

	curr, prev := head, &utils.ListNode{}
	dummy := prev
	for curr != nil {
		if _, ok := delSet[curr.Val]; !ok {
			prev.Next = curr
			prev = prev.Next
		} else {
			prev.Next = nil
		}

		curr = curr.Next
	}

	return dummy.Next
}

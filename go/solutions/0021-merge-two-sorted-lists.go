package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	result := &utils.ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}

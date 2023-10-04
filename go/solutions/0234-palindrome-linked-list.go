package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindromeList(head *utils.ListNode) bool {
	arr := make([]int, 0)
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	for i, j := 0, len(arr)-1; j > i; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}

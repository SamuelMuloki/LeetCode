package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func PairSum(head *utils.ListNode) int {
	curr, arr := head, []int{}
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	maxSum := 0
	for i := 0; i < len(arr)/2; i++ {
		maxSum = utils.Max(maxSum, arr[i]+arr[len(arr)-1-i])
	}

	return maxSum
}

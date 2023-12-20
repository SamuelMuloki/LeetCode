package solutions

import (
	"math/rand"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type RandomNodeSolution struct {
	nums []int
}

func RandomNodeConstructor(head *utils.ListNode) RandomNodeSolution {
	curr, nums := head, []int{}
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	return RandomNodeSolution{nums: nums}
}

func (this *RandomNodeSolution) GetRandom() int {
	j := rand.Intn(len(this.nums))
	return this.nums[j]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

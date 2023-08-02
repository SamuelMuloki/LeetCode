// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package leetcode

import "sort"

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	this.PopUntilK()
	return this.nums[0]
}

func (this *KthLargest) PopUntilK() {
	sort.Ints(this.nums)
	if len(this.nums) > this.k {
		this.nums = this.nums[len(this.nums)-this.k:]
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

package leetcode

/*
Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

E.g
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func ProductExceptSelf(nums []int) []int {
	prod := make([]int, len(nums))

	for i := range nums {
		prod[i] = 1
	}

	temp := 1
	for i := 0; i < len(nums); i++ {
		prod[i] = temp
		temp *= nums[i]
	}

	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod[i] *= temp
		temp *= nums[i]
	}

	return prod
}

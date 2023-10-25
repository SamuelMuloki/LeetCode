package solutions

func RotateArray(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse2(nums, 0, n-1)
	reverse2(nums, 0, k-1)
	reverse2(nums, k, n-1)
}

func reverse2(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

package solutions

func CheckPossibility(nums []int) bool {
	ans := 0
	for i := 1; i < len(nums) && ans <= 1; i++ {
		if nums[i-1] > nums[i] {
			ans++
			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}

	return ans <= 1
}

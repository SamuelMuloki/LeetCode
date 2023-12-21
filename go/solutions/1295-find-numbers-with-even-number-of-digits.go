package solutions

func FindNumbers(nums []int) int {
	ans := 0
	for i := range nums {
		if (nums[i] > 9 && nums[i] < 100) || (nums[i] > 999 && nums[i] < 10000) || nums[i] == 100000 {
			ans++
		}
	}

	return ans
}

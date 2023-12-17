package solutions

func FindNumbers(nums []int) int {
	ans := 0
	for i := range nums {
		count := 0
		for num := nums[i]; num > 0; num /= 10 {
			count++
		}

		if count%2 == 0 {
			ans++
		}
	}

	return ans
}

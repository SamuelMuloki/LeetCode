package solutions

func DifferenceOfSum(nums []int) int {
	elementSum, digitSum := 0, 0
	for i := range nums {
		elementSum += nums[i]
		num := nums[i]
		for num > 0 {
			digitSum += num % 10
			num /= 10
		}
	}

	return abs(elementSum - digitSum)
}

package solutions

func SingleNumber2(nums []int) int {
	ones, twos := 0, 0
	for i := range nums {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}

	return ones
}

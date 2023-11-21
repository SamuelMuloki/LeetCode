package solutions

func CountNicePairs(nums []int) int {
	rev := make([]int, len(nums))
	for i := range nums {
		reversed := 0
		for temp := nums[i]; temp > 0; temp /= 10 {
			reversed = (reversed * 10) + (temp % 10)
		}

		rev[i] = reversed - nums[i]
	}

	ans := 0
	set := make(map[int]int)
	for _, num := range rev {
		ans = (ans + set[num]) % (1e9 + 7)
		set[num]++
	}

	return ans
}

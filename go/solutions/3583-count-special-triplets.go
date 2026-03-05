package solutions

func SpecialTriplets(nums []int) int {
	const MOD = 1_000_000_007
	n := len(nums)

	right := make(map[int]int)
	for _, num := range nums {
		right[num]++
	}

	res := 0
	left := make(map[int]int)
	for j := 0; j < n; j++ {
		right[nums[j]]--
		if right[nums[j]] == 0 {
			delete(right, nums[j])
		}

		if j > 0 {
			target := nums[j] * 2
			leftCnt := left[target]
			rightCnt := right[target]

			res = (res + leftCnt*rightCnt) % MOD
		}
		left[nums[j]]++
	}

	return res
}

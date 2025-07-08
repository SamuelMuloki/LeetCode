package solutions

func MinimumOperations3(nums []int) int {
	cnt := [101]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		cnt[nums[i]]++
		if cnt[nums[i]] > 1 {
			return i/3 + 1
		}
	}

	return 0
}

package solutions

func MinimumOperations3(nums []int) int {
	cnt := [101]int{}
	lastRepeatedChar := 0
	for i := len(nums) - 1; i >= 0; i-- {
		cnt[nums[i]]++
		if cnt[nums[i]] > 1 {
			lastRepeatedChar = i + 1
			break
		}
	}

	div := lastRepeatedChar / 3
	if lastRepeatedChar%3 != 0 {
		return div + 1
	}

	return div
}

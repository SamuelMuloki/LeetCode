package solutions

func CountInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	var res int64
	currCount := 0
	freq := make(map[int]int)
	freq[0] = 1
	for end := 0; end < n; end++ {
		if nums[end]%modulo == k {
			currCount++
		}

		currMod := currCount % modulo
		target := (currMod - k + modulo) % modulo
		res += int64(freq[target])
		freq[currMod]++
	}

	return res
}

package solutions

func MaxFrequencyElements(nums []int) int {
	freq, maxFreq := make(map[int]int), 0
	for i := range nums {
		freq[nums[i]]++

		maxFreq = max(maxFreq, freq[nums[i]])
	}

	ans := 0
	for _, val := range freq {
		if val == maxFreq {
			ans += val
		}
	}

	return ans
}

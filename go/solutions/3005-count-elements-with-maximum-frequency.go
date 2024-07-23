package solutions

func MaxFrequencyElements(nums []int) int {
	freq, maxFreq, ans := make(map[int]int), 0, 0
	for i := range nums {
		freq[nums[i]]++

		frequency := freq[nums[i]]
		if frequency > maxFreq {
			maxFreq, ans = frequency, frequency
		} else if frequency == maxFreq {
			ans += frequency
		}
	}

	return ans
}

package solutions

func CountInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	cnt := make([]int, n+1)
	curr := 0
	for i, num := range nums {
		if num%modulo == k {
			curr++
		}
		cnt[i+1] = curr
	}

	var res int64
	freq := make(map[int]int)
	freq[0] = 1
	for end := 0; end < n; end++ {
		currSum := cnt[end+1] % modulo
		target := (currSum - k + modulo) % modulo
		res += int64(freq[target])
		freq[currSum]++
	}

	return res
}

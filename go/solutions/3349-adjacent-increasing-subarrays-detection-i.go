package solutions

func HasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if k <= 0 || 2*k > n {
		return false
	}

	good := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			good[i] = 1
		}
	}

	windowSize := k - 1
	pref := make([]int, len(good)+1)
	for i := 0; i < len(good); i++ {
		pref[i+1] = pref[i] + good[i]
	}

	for a := 0; a <= n-2*k; a++ {
		sum1 := pref[a+windowSize] - pref[a]
		if sum1 != windowSize {
			continue
		}
		b := a + k
		sum2 := pref[b+windowSize] - pref[b]
		if sum2 == windowSize {
			return true
		}
	}

	return false
}

package solutions

func MinKBitFlips(nums []int, k int) int {
	N := len(nums)
	flipped := make([]bool, N)
	validFlipsFromPastWindow, flipCount := 0, 0
	for i := 0; i < N; i++ {
		if i >= k {
			if flipped[i-k] {
				validFlipsFromPastWindow--
			}
		}

		if validFlipsFromPastWindow%2 == nums[i] {
			if i+k > N {
				return -1
			}

			validFlipsFromPastWindow++
			flipped[i] = true
			flipCount++
		}
	}

	return flipCount
}

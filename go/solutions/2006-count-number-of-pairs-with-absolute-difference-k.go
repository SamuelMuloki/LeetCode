package solutions

/*
Since nums[i] is in the [1...100] range, we can count each number using an array.
Now, we can sweep the counts, and accumulate the product of k-apart counts.
*/
func CountKDifference(nums []int, k int) int {
	ans, cnt := 0, make([]int, 101)
	for i := range nums {
		cnt[nums[i]]++
	}

	for i := k + 1; i < 101; i++ {
		ans += cnt[i] * cnt[i-k]
	}

	return ans
}

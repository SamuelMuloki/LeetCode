package solutions

func FindKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	next := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if nums[i] == key {
			next[i] = i
		} else if i+1 < n {
			next[i] = next[i+1]
		}
	}

	res := []int{}
	prev := -1
	for i := 0; i < n; i++ {
		if nums[i] == key {
			prev = i
		}
		if abs(next[i]-i) <= k || (prev != -1 && abs(prev-i) <= k) {
			res = append(res, i)
		}
	}

	return res
}

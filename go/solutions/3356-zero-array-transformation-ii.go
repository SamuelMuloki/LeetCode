package solutions

func MinZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	var canFormZeroArray = func(k int) bool {
		sum := 0
		diff := make([]int, n+1)
		for i := 0; i < k; i++ {
			diff[queries[i][0]] += queries[i][2]
			diff[queries[i][1]+1] -= queries[i][2]
		}

		for i := 0; i < n; i++ {
			sum += diff[i]
			if sum < nums[i] {
				return false
			}
		}

		return true
	}

	l, r := 0, len(queries)
	if !canFormZeroArray(r) {
		return -1
	}

	for l <= r {
		mid := l + (r-l)/2
		if canFormZeroArray(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

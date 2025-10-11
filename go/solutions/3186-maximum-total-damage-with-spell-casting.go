package solutions

import "sort"

func MaximumTotalDamage(power []int) int64 {
	freq := make(map[int]int)
	for _, num := range power {
		freq[num] += num
	}

	unique := make([]int, 0, len(freq))
	for k := range freq {
		unique = append(unique, k)
	}

	sort.Ints(unique)
	var binSearch = func(i, target int) int {
		lo, hi := 0, i-1
		res := -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if unique[mid] < target {
				res = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		return res
	}

	n := len(unique)
	dp := make([]int, n)
	dp[0] = freq[unique[0]]
	for i := 1; i < n; i++ {
		j := binSearch(i, unique[i]-2)
		if j == -1 {
			dp[i] = max(dp[i-1], freq[unique[i]])
		} else {
			dp[i] = max(dp[i-1], freq[unique[i]]+dp[j])
		}
	}

	return int64(dp[n-1])
}

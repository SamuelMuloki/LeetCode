package solutions

import "sort"

func MaxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	var extra int64 = 0
	m := len(batteries)
	for i := 0; i < m-n; i++ {
		extra += int64(batteries[i])
	}

	live := batteries[m-n:]

	for i := 0; i < n-1; i++ {
		diff := int64(live[i+1] - live[i])
		need := int64(i+1) * diff
		if extra < need {
			return int64(live[i]) + extra/int64(i+1)
		}
		extra -= need
	}

	return int64(live[n-1]) + extra/int64(n)
}

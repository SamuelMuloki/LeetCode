package solutions

import "sort"

func MaximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	ans := 0
	for i := 0; i < k; i++ {
		if happiness[i]-i > 0 {
			ans += happiness[i] - i
		}
	}

	return int64(ans)
}

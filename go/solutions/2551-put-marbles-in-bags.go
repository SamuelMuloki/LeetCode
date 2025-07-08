package solutions

import "sort"

func PutMarbles(weights []int, k int) int64 {
	n := len(weights)

	if k == n || k == 1 {
		return 0
	}

	pairSums := make([]int, 0, n-1)
	for i := 0; i < n-1; i++ {
		pairSums = append(pairSums, weights[i]+weights[i+1])
	}

	sort.Ints(pairSums)

	minScore := int64(0)
	maxScore := int64(0)

	for i := 0; i < k-1; i++ {
		minScore += int64(pairSums[i])
		maxScore += int64(pairSums[n-2-i])
	}

	return maxScore - minScore
}

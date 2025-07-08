package solutions

import "sort"

func MaxValue(events [][]int, k int) int {
	n := len(events)
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	for curIndex := n - 1; curIndex >= 0; curIndex-- {
		nextIndex := sort.Search(len(events), func(i int) bool {
			return events[i][0] > events[curIndex][1]
		})
		for count := 1; count <= k; count++ {
			dp[count][curIndex] = max(dp[count][curIndex+1], events[curIndex][2]+dp[count-1][nextIndex])
		}
	}

	return dp[k][0]
}

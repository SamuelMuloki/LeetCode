package solutions

func MinHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	dp[1] = books[0][1]

	for i := 2; i <= len(books); i++ {
		remainingShelfWidth := shelfWidth - books[i-1][0]
		maxHeight := books[i-1][1]
		dp[i] = dp[i-1] + maxHeight

		j := i - 1
		for j > 0 && remainingShelfWidth-books[j-1][0] >= 0 {
			remainingShelfWidth -= books[j-1][0]
			maxHeight = max(maxHeight, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+maxHeight)
			j--
		}
	}

	return dp[n]
}

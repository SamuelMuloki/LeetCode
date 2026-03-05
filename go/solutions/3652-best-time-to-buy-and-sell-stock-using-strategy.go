package solutions

func MaxProfit4(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// Calculate initial profit without any modification
	initialProfit := int64(0)
	for i := 0; i < n; i++ {
		initialProfit += int64(strategy[i]) * int64(prices[i])
	}

	// Precompute prefix sums for prices and strategy*prices
	pricePrefix := make([]int64, n+1)
	strategyPricePrefix := make([]int64, n+1)

	for i := 0; i < n; i++ {
		pricePrefix[i+1] = pricePrefix[i] + int64(prices[i])
		strategyPricePrefix[i+1] = strategyPricePrefix[i] + int64(strategy[i])*int64(prices[i])
	}

	half := k / 2
	maxDelta := int64(0)

	// Slide the window and calculate delta for each position
	for start := 0; start <= n-k; start++ {
		end := start + k

		// New contribution: sum of prices in the second half of the window
		newContribution := pricePrefix[end] - pricePrefix[start+half]

		// Old contribution: sum of strategy[i]*prices[i] in the window
		oldContribution := strategyPricePrefix[end] - strategyPricePrefix[start]

		// Delta is the change in profit
		delta := newContribution - oldContribution

		// Track maximum delta
		if delta > maxDelta {
			maxDelta = delta
		}
	}

	return initialProfit + maxDelta
}

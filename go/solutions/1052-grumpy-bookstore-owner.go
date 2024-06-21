package solutions

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	n, unrealizedCustomers := len(customers), 0

	for i := 0; i < minutes; i++ {
		unrealizedCustomers += customers[i] * grumpy[i]
	}

	maxUnrealizedCustomers := unrealizedCustomers
	for i := minutes; i < n; i++ {
		unrealizedCustomers += customers[i] * grumpy[i]
		unrealizedCustomers -= customers[i-minutes] * grumpy[i-minutes]
		maxUnrealizedCustomers = max(maxUnrealizedCustomers, unrealizedCustomers)
	}

	totalCustomers := maxUnrealizedCustomers
	for i := 0; i < n; i++ {
		totalCustomers += customers[i] * (1 - grumpy[i])
	}

	return totalCustomers
}

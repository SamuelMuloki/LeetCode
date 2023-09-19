package solutions

const MOD = 1000000007

func CountOrders(n int) int {
	count := 1
	for i := 2; i <= n; i++ {
		count = (count * (2*i - 1) * i) % MOD
	}
	return count
}

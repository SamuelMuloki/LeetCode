package solutions

func CountGoodNumbers(n int64) int {
	var MOD int64 = 1e9 + 7
	var binpow = func(a, b int64) int64 {
		var res int64 = 1
		for b > 0 {
			if b&1 == 1 {
				res = (res * a) % MOD
			}
			a = (a * a) % MOD
			b >>= 1
		}

		return res
	}

	var even int64 = (n + 1) / 2
	var odd int64 = n / 2
	return int((binpow(5, even) * binpow(4, odd)) % MOD)
}

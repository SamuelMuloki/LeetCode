package solutions

// sieve of eratosthenes
func CountPrimes(n int) int {
	if n == 0 {
		return 0
	}

	//all elements marked prime initially
	prime := make([]bool, n)
	for i := 2; i < len(prime); i++ {
		prime[i] = true
	}

	count := 0
	for i := 2; i < n; i++ {
		//if the element is prime then mark all its multiples as not prime
		if prime[i] {
			count++
			j := i * 2
			for j < n {
				prime[j] = false
				j += i
			}
		}
	}

	return count
}

package solutions

import "math"

func PrimeSubOperation(nums []int) bool {
	var isPrime = func(num int) bool {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	maxElement := math.MinInt
	for _, num := range nums {
		maxElement = max(maxElement, num)
	}

	previousPrime := make([]int, maxElement+1)
	for i := 2; i <= maxElement; i++ {
		if isPrime(i) {
			previousPrime[i] = i
		} else {
			previousPrime[i] = previousPrime[i-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		bound := 0
		if i == 0 {
			bound = nums[0]
		} else {
			bound = nums[i] - nums[i-1]
		}

		if bound <= 0 {
			return false
		}

		largestPrime := previousPrime[bound-1]
		nums[i] -= largestPrime
	}

	return true
}

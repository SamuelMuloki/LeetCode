package solutions

import (
	"math"
	"sort"
)

func MaximumScore2(nums []int, k int) int {
	const MOD = 1e9 + 7
	n := len(nums)

	getPrimes := func(limit int) []int {
		isPrime := make([]bool, limit+1)
		for i := range isPrime {
			isPrime[i] = true
		}
		primes := []int{}
		for number := 2; number <= limit; number++ {
			if !isPrime[number] {
				continue
			}
			primes = append(primes, number)
			for multiple := number * number; multiple <= limit; multiple += number {
				isPrime[multiple] = false
			}
		}
		return primes
	}

	power := func(base, exponent int) int {
		res := 1
		for exponent > 0 {
			if exponent%2 == 1 {
				res = (res * base) % MOD
			}
			base = (base * base) % MOD
			exponent /= 2
		}
		return res
	}

	maxElement := 0
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}

	primes := getPrimes(maxElement)

	primeScores := make([]int, n)
	for i, num := range nums {
		for _, prime := range primes {
			if prime*prime > num {
				break
			}
			if num%prime != 0 {
				continue
			}
			primeScores[i]++
			for num%prime == 0 {
				num /= prime
			}
		}
		if num > 1 {
			primeScores[i]++
		}
	}

	nextDominant := make([]int, n)
	prevDominant := make([]int, n)
	for i := range nextDominant {
		nextDominant[i] = n
		prevDominant[i] = -1
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && primeScores[stack[len(stack)-1]] < primeScores[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextDominant[top] = i
		}
		if len(stack) > 0 {
			prevDominant[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	numOfSubarrays := make([]int64, n)
	for i := 0; i < n; i++ {
		numOfSubarrays[i] = int64(nextDominant[i]-i) * int64(i-prevDominant[i])
	}

	sortedArray := make([][2]int, n)
	for i, num := range nums {
		sortedArray[i] = [2]int{num, i}
	}
	sort.Slice(sortedArray, func(i, j int) bool {
		return sortedArray[i][0] > sortedArray[j][0]
	})

	score := 1
	for _, elem := range sortedArray {
		num, index := elem[0], elem[1]
		operations := int(math.Min(float64(k), float64(numOfSubarrays[index])))
		score = (score * power(num, operations)) % MOD
		k -= operations
		if k == 0 {
			break
		}
	}

	return score
}

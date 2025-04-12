package solutions

import (
	"fmt"
	"strconv"
)

func CountGoodIntegers(n int, k int) int64 {
	palindromes := generateNDigitPalindromes(n)
	seenMultisets := make(map[[10]int]int64)

	for _, pal := range palindromes {
		if pal%k != 0 {
			continue
		}

		digitFreq := [10]int{}
		numStr := fmt.Sprintf("%0*d", n, pal)
		for _, c := range numStr {
			digit := int(c - '0')
			digitFreq[digit]++
		}

		if _, exists := seenMultisets[digitFreq]; !exists {
			count := countPermutations(digitFreq[:], n)
			seenMultisets[digitFreq] = count
		}
	}

	total := int64(0)
	for _, count := range seenMultisets {
		total += count
	}

	return total
}

func factorial(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func countPermutations(digitFreq []int, n int) int64 {
	totalValid := int64(0)
	for firstDigit := 1; firstDigit <= 9; firstDigit++ {
		if digitFreq[firstDigit] == 0 {
			continue
		}
		freqCopy := make([]int, 10)
		copy(freqCopy, digitFreq)
		freqCopy[firstDigit]--
		numerator := factorial(n - 1)
		denominator := int64(1)
		for _, freq := range freqCopy {
			denominator *= factorial(freq)
		}
		count := numerator / denominator
		totalValid += count
	}
	return totalValid
}

func generateNDigitPalindromes(n int) []int {
	var palindromes []int

	if n == 1 {
		for i := 1; i <= 9; i++ {
			palindromes = append(palindromes, i)
		}
		return palindromes
	}

	halfDigits := (n + 1) / 2
	startHalf := 1
	for i := 1; i < halfDigits; i++ {
		startHalf *= 10
	}
	endHalf := startHalf*10 - 1

	for i := startHalf; i <= endHalf; i++ {
		firstHalf := strconv.Itoa(i)
		var palindromeStr string
		if n%2 == 0 {
			reverse := reverseString(firstHalf)
			palindromeStr = firstHalf + reverse
		} else {
			reverse := reverseString(firstHalf[:len(firstHalf)-1])
			palindromeStr = firstHalf + reverse
		}

		num, _ := strconv.Atoi(palindromeStr)
		palindromes = append(palindromes, num)
	}

	return palindromes
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

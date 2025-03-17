package solutions

import "math/bits"

func MinimizeXor(num1 int, num2 int) int {
	ans := 0
	targetSetBitCount := bits.OnesCount(uint(num2))
	setBitsCount, currentBit := 0, 31

	var isSet = func(x, bit int) bool {
		return (x & (1 << bit)) != 0
	}

	var setBit = func(x, bit int) int {
		return x | (1 << bit)
	}

	for setBitsCount < targetSetBitCount {
		if isSet(num1, currentBit) ||
			targetSetBitCount-setBitsCount > currentBit {
			ans = setBit(ans, currentBit)
			setBitsCount++
		}
		currentBit--
	}

	return ans
}

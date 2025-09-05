package solutions

import "math/bits"

func MakeTheIntegerZero(num1 int, num2 int) int {
	k := 1
	for {
		x := int64(num1) - int64(num2)*int64(k)
		if x < int64(k) {
			return -1
		}
		if k >= bits.OnesCount(uint(x)) {
			return k
		}
		k++
	}
}

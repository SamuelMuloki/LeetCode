package solutions

func NthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	var i2, i3, i5 int
	ugly := make([]int, n)
	ugly[0] = 1

	for i := 1; i < n; i++ {
		ugly[i] = min(ugly[i2]*2, min(ugly[i3]*3, ugly[i5]*5))

		if ugly[i] == ugly[i2]*2 {
			i2++
		}
		if ugly[i] == ugly[i3]*3 {
			i3++
		}
		if ugly[i] == ugly[i5]*5 {
			i5++
		}
	}

	return ugly[n-1]
}

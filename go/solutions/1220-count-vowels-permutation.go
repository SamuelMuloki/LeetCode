package solutions

func CountVowelPermutation(n int) int {
	mod := 1000000007
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 1; j < n; j++ {
		aNext, eNext, iNext := e, (a+i)%mod, (a+e+o+u)%mod
		oNext, uNext := (i+u)%mod, a

		a, e, i, o, u = aNext, eNext, iNext, oNext, uNext
	}

	return (a + e + i + o + u) % mod
}

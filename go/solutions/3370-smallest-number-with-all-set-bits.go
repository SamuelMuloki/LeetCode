package solutions

func SmallestNumber2(n int) int {
	ans := 1
	for c := 1; ans < n; c++ {
		ans += 1 << c
	}

	return ans
}

package solutions

func MinOperations12(nums []int) int {
	n := len(nums)
	num1 := 0
	g := 0
	for _, x := range nums {
		if x == 1 {
			num1++
		}
		g = gcd(g, x)
	}
	if num1 > 0 {
		return n - num1
	}
	if g > 1 {
		return -1
	}

	minLen := n
	for i := 0; i < n; i++ {
		currentGcd := 0
		for j := i; j < n; j++ {
			currentGcd = gcd(currentGcd, nums[j])
			if currentGcd == 1 {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}
	return minLen + n - 2
}

package solutions

// https://leetcode.com/problems/xor-operation-in-an-array/solutions/699141/visual-solution-python-o-1-time-o-1-space/
func XorOperation(n int, start int) int {
	last := start + 2*(n-1)
	if start%4 <= 1 {
		if n%4 == 1 {
			return last
		} else if n%4 == 2 {
			return 2
		} else if n%4 == 3 {
			return 2 ^ last
		} else {
			return 0
		}
	}

	if n%4 == 1 {
		return start
	} else if n%4 == 2 {
		return start ^ last
	} else if n%4 == 3 {
		return start ^ 2
	}

	return start ^ 2 ^ last
}

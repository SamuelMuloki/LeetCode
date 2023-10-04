package solutions

func MinNumberOperations(target []int) int {
	res := target[0]
	for i := 1; i < len(target); i++ {
		prev, curr := target[i-1], target[i]
		if curr > prev {
			res += curr - prev
		}
	}

	return res
}

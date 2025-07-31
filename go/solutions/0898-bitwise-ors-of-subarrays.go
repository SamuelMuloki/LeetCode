package solutions

func SubarrayBitwiseORs(arr []int) int {
	res := make(map[int]bool)
	curr := make(map[int]bool)

	for _, num := range arr {
		next_curr := map[int]bool{num: true}
		res[num] = true
		for prev := range curr {
			next_curr[prev|num] = true
			res[prev|num] = true
		}
		curr = next_curr
	}

	return len(res)
}

package solutions

func MaxOperations(nums []int, k int) int {
	var operations int
	complements := make(map[int]int)

	for _, num := range nums {
		comp := k - num
		if complements[comp] != 0 {
			operations++
			complements[comp]--
		} else {
			complements[num]++
		}
	}

	return operations
}

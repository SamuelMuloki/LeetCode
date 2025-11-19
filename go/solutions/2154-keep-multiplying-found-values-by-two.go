package solutions

func FindFinalValue(nums []int, original int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	for set[original] {
		original = 2 * original
	}

	return original
}

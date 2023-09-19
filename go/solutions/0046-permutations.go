package solutions

func Permute(nums []int) [][]int {
	output := make([][]int, 0)

	var backtrack func(curr []int, l int, r int)
	backtrack = func(curr []int, l int, r int) {
		if l == r {
			output = append(output, append([]int{}, curr...))
		}

		for i := l; i <= r; i++ {
			curr[l], curr[i] = curr[i], curr[l]
			backtrack(curr, l+1, r)
			curr[l], curr[i] = curr[i], curr[l]
		}
	}

	backtrack(nums, 0, len(nums)-1)

	return output
}

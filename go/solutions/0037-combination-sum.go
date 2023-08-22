package solutions

func CombinationSum(candidates []int, target int) [][]int {
	output := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(idx int, currSum int, curr []int)
	backtrack = func(idx int, currSum int, curr []int) {
		if currSum == target {
			output = append(output, append([]int{}, curr...))
			return
		}

		if currSum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(i, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, 0, curr)

	return output
}

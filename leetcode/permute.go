package leetcode

func Permute(nums []int) [][]int {
	output := make([][]int, 0)
	curr := make([]int, 0)
	n := len(nums)
	visited := make(map[int]int)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(curr) == n {
			output = append(output, append([]int{}, curr...))
		}

		for i := 0; i < n; i++ {
			if visited[i] == 0 {
				visited[i]++
				curr = append(curr, nums[i])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
				visited[i]--
			}
		}
	}

	backtrack(0)

	return output
}

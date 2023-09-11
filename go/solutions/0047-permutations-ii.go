package solutions

import "fmt"

func PermuteUnique(nums []int) [][]int {
	output := make([][]int, 0)

	var backtrack func(curr []int, l int, r int)
	backtrack = func(curr []int, l int, r int) {
		if l == r {
			output = append(output, append([]int{}, curr...))
		}

		visited := make(map[int]int)
		for i := l; i <= r; i++ {
			fmt.Println(visited, curr, i, l)
			if visited[curr[i]] == 0 {
				curr[l], curr[i] = curr[i], curr[l]
				backtrack(curr, l+1, r)
				curr[l], curr[i] = curr[i], curr[l]
			}
			visited[curr[i]]++
		}
	}

	backtrack(nums, 0, len(nums)-1)

	return output
}

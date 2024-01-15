package solutions

func FindWinners(matches [][]int) [][]int {
	scores := [100001]int{}
	for i := range matches {
		if scores[matches[i][0]] == 0 {
			scores[matches[i][0]] = 1
		}

		if scores[matches[i][1]] == 0 {
			scores[matches[i][1]] = 1
		}
		scores[matches[i][1]]++
	}

	zeros, ones := make([]int, 0), make([]int, 0)
	for i := range scores {
		if scores[i] == 1 {
			zeros = append(zeros, i)
		}

		if scores[i] == 2 {
			ones = append(ones, i)
		}
	}

	return [][]int{zeros, ones}
}

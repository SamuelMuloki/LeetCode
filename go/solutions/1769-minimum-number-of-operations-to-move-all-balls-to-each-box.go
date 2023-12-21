package solutions

func MinOperations3(boxes string) []int {
	ans := make([]int, len(boxes))
	for i := range boxes {
		count := 0
		for j := 0; j < len(boxes); j++ {
			if i != j && boxes[j] == '1' {
				count += abs(j - i)
			}
		}

		ans[i] = count
	}

	return ans
}

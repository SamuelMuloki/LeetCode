package solutions

func NumberOfBeams(bank []string) int {
	ans, prev := 0, 0
	for i := range bank {
		count := 0
		for j := range bank[i] {
			if bank[i][j] == '1' {
				count++
			}
		}
		if count != 0 {
			ans += count * prev
			prev = count
		}
	}

	return ans
}

package solutions

func WinnerOfGame(colors string) bool {
	aCount, bCount := 0, 0
	for i := 1; i < len(colors)-1; i++ {
		if colors[i-1] == colors[i] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				aCount++
			} else {
				bCount++
			}
		}
	}

	return aCount > bCount
}

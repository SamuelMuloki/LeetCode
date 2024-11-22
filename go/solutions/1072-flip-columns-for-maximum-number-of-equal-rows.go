package solutions

func MaxEqualRowsAfterFlips(matrix [][]int) int {
	patternFrequency := make(map[string]int)
	for _, row := range matrix {
		pattern := ""
		for col := 0; col < len(row); col++ {
			if row[0] == row[col] {
				pattern += "T"
			} else {
				pattern += "F"
			}
		}
		patternFrequency[pattern]++
	}

	maxFrequency := 0
	for _, freq := range patternFrequency {
		maxFrequency = max(maxFrequency, freq)
	}

	return maxFrequency
}

package solutions

func StrangePrinter(s string) int {
	var removeDuplicates = func(s string) string {
		var uniqueChars string
		i := 0
		for i < len(s) {
			currentChar := s[i]
			uniqueChars += string(currentChar)
			for i < len(s) && s[i] == currentChar {
				i++
			}
		}

		return uniqueChars
	}

	str := removeDuplicates(s)
	n := len(str)
	minTurns := make([][]int, n)
	for i := range minTurns {
		minTurns[i] = make([]int, n)
		minTurns[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for start := 0; start+length-1 < n; start++ {
			end := start + length - 1
			minTurns[start][end] = length

			for split := 0; split < length-1; split++ {
				totalTurns := minTurns[start][start+split] + minTurns[start+split+1][end]

				if str[start+split] == str[end] {
					totalTurns--
				}

				minTurns[start][end] = min(minTurns[start][end], totalTurns)
			}
		}
	}

	return minTurns[0][n-1]
}

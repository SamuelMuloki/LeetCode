package solutions

func NumberOfWays2(corridor string) int {
	count, seats := 1, 0
	previousPairLast := -1
	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			seats++

			if seats == 2 {
				previousPairLast = i
				seats = 0
			} else if seats == 1 && previousPairLast != -1 {
				count *= (i - previousPairLast)
				count %= 1e9 + 7
			}
		}
	}

	if seats == 1 || previousPairLast == -1 {
		return 0
	}

	return count
}

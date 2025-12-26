package solutions

func BestClosingTime(customers string) int {
	curPenalty := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			curPenalty++
		}
	}

	minPenalty := curPenalty
	earliestHour := 0

	for i := 0; i < len(customers); i++ {
		ch := customers[i]

		if ch == 'Y' {
			curPenalty--
		} else {
			curPenalty++
		}

		if curPenalty < minPenalty {
			minPenalty = curPenalty
			earliestHour = i + 1
		}
	}

	return earliestHour
}

package solutions

func NumberOfArrays(differences []int, lower int, upper int) int {
	prefSum := 0
	minPref, maxPref := 0, 0
	for _, diff := range differences {
		prefSum += diff
		if prefSum < minPref {
			minPref = prefSum
		}
		if prefSum > maxPref {
			maxPref = prefSum
		}
	}

	return max(0, (upper-lower+1)-(maxPref-minPref))
}

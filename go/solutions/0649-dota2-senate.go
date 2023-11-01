package solutions

func PredictPartyVictory(senate string) string {
	rad, dir := []int{}, []int{}
	n := len(senate)
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			rad = append(rad, i)
		} else {
			dir = append(dir, i)
		}
	}

	for len(rad) > 0 && len(dir) > 0 {
		if rad[0] < dir[0] {
			rad = append(rad, n)
		} else {
			dir = append(dir, n)
		}

		rad = rad[1:]
		dir = dir[1:]
		n++
	}

	if len(rad) == 0 {
		return "Dire"
	}

	return "Radiant"
}

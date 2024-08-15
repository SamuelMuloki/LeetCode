package solutions

func LemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := range bills {
		if bills[i] == 5 {
			five++
		} else if five == 0 {
			return false
		} else if bills[i] == 10 {
			five--
			ten++
		} else if ten == 0 && five < 3 {
			return false
		} else {
			if ten == 0 && five >= 3 {
				five -= 3
			} else {
				ten--
				five--
			}
		}
	}

	return true
}

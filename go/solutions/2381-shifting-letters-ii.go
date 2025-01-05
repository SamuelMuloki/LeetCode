package solutions

func ShiftingLetters2(s string, shifts [][]int) string {
	n := len(s)
	diffArray := make([]int, n)
	for _, shift := range shifts {
		if shift[2] == 1 {
			diffArray[shift[0]]++

			if shift[1]+1 < n {
				diffArray[shift[1]+1]--
			}
		} else {
			diffArray[shift[0]]--

			if shift[1]+1 < n {
				diffArray[shift[1]+1]++
			}
		}
	}

	runes := []byte(s)
	numberOfShifts := 0
	for i := 0; i < n; i++ {
		numberOfShifts = (numberOfShifts + diffArray[i]) % 26
		if numberOfShifts < 0 {
			numberOfShifts += 26
		}
		runes[i] = byte((int(runes[i]-'a')+numberOfShifts)%26) + 'a'
	}

	return string(runes)
}

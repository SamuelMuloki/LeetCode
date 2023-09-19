package solutions

func NextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1

	for l < r {
		mid := l + (r-l)/2

		if letters[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if letters[l] <= target {
		return letters[0]
	}

	return letters[l]
}

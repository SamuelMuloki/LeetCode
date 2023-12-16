package solutions

func CheckPowersOfThree(n int) bool {
	if n == 1 {
		return true
	}

	if n%3 == 2 {
		return false
	}

	return CheckPowersOfThree(n / 3)
}

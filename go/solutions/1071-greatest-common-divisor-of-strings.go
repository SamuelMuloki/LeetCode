package solutions

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	var gcd func(x, y int) int
	gcd = func(x, y int) int {
		if y == 0 {
			return x
		} else {
			return gcd(y, x%y)
		}
	}

	gcdLen := gcd(len(str1), len(str2))

	return str1[:gcdLen]
}

package solutions

import "strconv"

func FractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	res := ""
	if (numerator < 0) != (denominator < 0) {
		res += "-"
	}
	n, d := abs(numerator), abs(denominator)
	res += strconv.Itoa(n / d)
	rem := n % d
	if rem == 0 {
		return res
	}
	res += "."
	m := map[int]int{}
	for rem != 0 {
		if pos, ok := m[rem]; ok {
			res = res[:pos] + "(" + res[pos:] + ")"
			break
		}
		m[rem] = len(res)
		rem *= 10
		res += strconv.Itoa(rem / d)
		rem %= d
	}
	return res
}

package solutions

import "strconv"

func AddBinary(a string, b string) string {
	ans, carry := "", 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}

		if j >= 0 {
			sum += int(b[j] - '0')
		}

		ans = strconv.Itoa(sum%2) + ans
		if sum <= 1 {
			carry = 0
		} else {
			carry = 1
		}
	}

	if carry == 1 {
		ans = strconv.Itoa(carry) + ans
	}

	return ans
}

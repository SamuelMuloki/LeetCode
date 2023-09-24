package solutions

func MaximumOddBinaryNumber(s string) string {
	count := 0
	for i := 0; i < len(s); i++ {
		count += int(s[i]&'1') - '0'
	}

	str := "1"
	r := len(s) - count
	for r > 0 {
		str = "0" + str
		r--
	}

	for count > 1 {
		str = "1" + str
		count--
	}

	return str
}

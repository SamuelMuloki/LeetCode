package solutions

func NumSteps(s string) int {
	operations, carry := 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if (int(s[i]-'0')+carry)%2 != 0 {
			operations += 2
			carry = 1
		} else {
			operations++
		}
	}

	return operations + carry
}

package solutions

func DivideString(s string, k int, fill byte) []string {
	res := []string{}
	curr := []byte{}
	for i := 0; i < len(s); i++ {
		curr = append(curr, s[i])
		if (i+1)%k == 0 {
			res = append(res, string(curr))
			curr = []byte{}
		}
	}

	if len(curr) != 0 {
		for x := len(curr); x < k; x++ {
			curr = append(curr, fill)
		}
		res = append(res, string(curr))
	}

	return res
}

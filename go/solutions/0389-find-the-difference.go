package solutions

func FindTheDifference(s string, t string) byte {
	s += t
	c := byte(0)
	for i := range s {
		c ^= s[i]
	}

	return c
}

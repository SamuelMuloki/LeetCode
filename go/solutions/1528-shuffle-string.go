package solutions

func RestoreString(s string, indices []int) string {
	runes := make([]byte, len(indices))
	for i := range indices {
		runes[indices[i]] = s[i]
	}

	return string(runes)
}

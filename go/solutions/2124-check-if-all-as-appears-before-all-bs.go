package solutions

func CheckString(s string) bool {
	foundB := false
	for i := range s {
		if s[i] == 'b' {
			foundB = true
		}
		if s[i] == 'a' && foundB {
			return false
		}
	}

	return true
}

package solutions

func RotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		arr := []byte(s[i:])
		arr = append(arr, s[:i]...)
		if string(arr) == string(goal) {
			return true
		}
	}

	return false
}

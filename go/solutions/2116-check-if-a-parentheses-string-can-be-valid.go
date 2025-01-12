package solutions

func CanBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	openBrackets := []int{}
	unlocked := []int{}

	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			unlocked = append(unlocked, i)
		} else if s[i] == '(' {
			openBrackets = append(openBrackets, i)
		} else if s[i] == ')' {
			if len(openBrackets) > 0 {
				openBrackets = openBrackets[:len(openBrackets)-1]
			} else if len(unlocked) > 0 {
				unlocked = unlocked[:len(unlocked)-1]
			} else {
				return false
			}
		}
	}

	for len(openBrackets) > 0 &&
		len(unlocked) > 0 &&
		openBrackets[len(openBrackets)-1] < unlocked[len(unlocked)-1] {
		openBrackets = openBrackets[:len(openBrackets)-1]
		unlocked = unlocked[:len(unlocked)-1]
	}

	if len(openBrackets) > 0 {
		return false
	}

	return true
}

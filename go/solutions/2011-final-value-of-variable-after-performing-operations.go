package solutions

func FinalValueAfterOperations(operations []string) int {
	ans := 0
	for i := range operations {
		if operations[i][1] == '+' {
			ans++
		} else {
			ans--
		}
	}

	return ans
}

package solutions

func FinalValueAfterOperations(operations []string) int {
	ans := 0
	for i := range operations {
		switch operations[i] {
		case "++X":
			ans++
		case "X++":
			ans++
		case "--X":
			ans--
		case "X--":
			ans--
		}
	}

	return ans
}

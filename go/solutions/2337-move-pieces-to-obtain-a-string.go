package solutions

func CanChange(start string, target string) bool {
	startLength := len(start)
	startIndex, targetIndex := 0, 0

	for startIndex < startLength || targetIndex < startLength {
		for startIndex < startLength && start[startIndex] == '_' {
			startIndex++
		}

		for targetIndex < startLength && target[targetIndex] == '_' {
			targetIndex++
		}

		if startIndex == startLength || targetIndex == startLength {
			return startIndex == startLength && targetIndex == startLength
		}

		if start[startIndex] != target[targetIndex] ||
			(start[startIndex] == 'L' && startIndex < targetIndex) ||
			(start[startIndex] == 'R' && startIndex > targetIndex) {
			return false
		}

		startIndex++
		targetIndex++
	}

	return true
}

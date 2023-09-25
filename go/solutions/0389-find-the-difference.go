package solutions

func FindTheDifference(s string, t string) byte {
	newArr := []byte(t)
	for i := range s {
		newArr[i+1] += newArr[i] - s[i]
	}

	return newArr[len(newArr)-1]
}

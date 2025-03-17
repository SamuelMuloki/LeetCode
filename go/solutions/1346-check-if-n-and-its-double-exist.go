package solutions

func CheckIfExist(arr []int) bool {
	count := make(map[int]int)
	for i := range arr {
		count[2*arr[i]] = i
	}

	for i := range arr {
		if idx, ok := count[arr[i]]; ok && idx != i {
			return true
		}
	}

	return false
}

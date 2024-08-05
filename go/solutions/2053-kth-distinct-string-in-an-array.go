package solutions

func KthDistinct(arr []string, k int) string {
	count := make(map[string]int)
	for i := range arr {
		count[arr[i]]++
	}

	for i := range arr {
		if count[arr[i]] > 1 {
			continue
		}
		k--
		if k == 0 {
			return arr[i]
		}
	}

	return ""
}

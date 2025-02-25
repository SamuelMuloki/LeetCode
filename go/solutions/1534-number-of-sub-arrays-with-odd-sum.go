package solutions

func NumOfSubarrays(arr []int) int {
	count, prefSum := 0, 0
	oddCount, evenCount := 0, 1
	for _, num := range arr {
		prefSum += num
		if prefSum%2 == 0 {
			count += oddCount
			evenCount++
		} else {
			count += evenCount
			oddCount++
		}

		count %= (1e9 + 7)
	}

	return count
}

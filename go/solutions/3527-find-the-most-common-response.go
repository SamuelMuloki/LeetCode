package solutions

func FindCommonResponse(responses [][]string) string {
	cnt := make(map[string]int)
	maxCnt := 0
	for _, response := range responses {
		set := make(map[string]bool)
		for _, word := range response {
			if !set[word] {
				cnt[word]++
				maxCnt = max(maxCnt, cnt[word])
			}
			set[word] = true
		}
	}

	strs := []string{}
	for word, val := range cnt {
		if val == maxCnt {
			strs = append(strs, word)
		}
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return lexicographicallySmallestString(strs)
}

func lexicographicallySmallestString(arr []string) string {
	res := arr[0]
	for _, str := range arr {
		if str < res {
			res = str
		}
	}

	return res
}

package solutions

func MaxLength(arr []string) int {
	res := make([]string, 0)
	res = append(res, "")
	for _, str := range arr {
		if !isUnique(str) {
			continue
		}

		resArr := make([]string, 0)
		for i := range res {
			temp := res[i] + str
			if isUnique(temp) {
				resArr = append(resArr, temp)
			}
		}
		res = append(res, resArr...)
	}

	ans := 0
	for _, str := range res {
		ans = max(ans, len(str))
	}

	return ans
}

func isUnique(str string) bool {
	if len(str) > 26 {
		return false
	}

	used := [26]bool{}
	for _, ch := range str {
		if used[ch-'a'] {
			return false
		} else {
			used[ch-'a'] = true
		}
	}

	return true
}

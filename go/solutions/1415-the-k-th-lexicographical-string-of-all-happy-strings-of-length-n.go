package solutions

func GetHappyString(n int, k int) string {
	letters := []string{"a", "b", "c"}
	arr := []string{}
	var findString func(curr string)
	findString = func(curr string) {
		if len(arr) == k {
			return
		}

		if len(curr) == n {
			arr = append(arr, curr)
			return
		}

		for i := 0; i < len(letters); i++ {
			if len(curr) > 0 && string(curr[len(curr)-1]) == letters[i] {
				continue
			}
			curr += string(letters[i])
			findString(curr)
			curr = curr[:len(curr)-1]
		}
	}

	findString("")

	if k > len(arr) {
		return ""
	}

	return arr[len(arr)-1]
}

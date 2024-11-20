package solutions

func TakeCharacters(s string, k int) int {
	count := [3]int{}
	n := len(s)

	for _, ch := range s {
		count[ch-'a']++
	}

	for i := 0; i < 3; i++ {
		if count[i] < k {
			return -1
		}
	}

	window := [3]int{}
	left, maxWindow := 0, 0
	for right := 0; right < n; right++ {
		window[s[right]-'a']++

		for left <= right && (count[0]-window[0] < k || count[1]-window[1] < k || count[2]-window[2] < k) {
			window[s[left]-'a']--
			left++
		}

		maxWindow = max(maxWindow, right-left+1)
	}

	return n - maxWindow
}

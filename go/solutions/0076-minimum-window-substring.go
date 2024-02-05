package solutions

func MinWindow(s string, t string) string {
	var ans string
	cache := [128]int{}
	for i := range t {
		cache[t[i]]++
	}

	cnt := len(t)
	for left, right := 0, 0; right < len(s); right++ {
		if cache[s[right]] > 0 {
			cnt--
		}
		cache[s[right]]--
		for cnt == 0 {
			if ans == "" || right-left < len(ans) {
				ans = s[left : right+1]
			}
			cache[s[left]]++
			if cache[s[left]] > 0 {
				cnt++
			}
			left++
		}
	}
	return ans
}

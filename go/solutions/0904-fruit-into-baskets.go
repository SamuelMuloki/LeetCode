package solutions

func TotalFruit(fruits []int) int {
	n := len(fruits)
	set := make(map[int]int)
	start := 0
	res := 0
	for end := 0; end < n; end++ {
		set[fruits[end]]++
		for len(set) > 2 {
			set[fruits[start]]--
			if set[fruits[start]] == 0 {
				delete(set, fruits[start])
			}
			start++
		}
		res = max(res, end-start+1)
	}

	return res
}

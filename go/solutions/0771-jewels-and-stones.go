package solutions

func NumJewelsInStones(jewels string, stones string) int {
	set := [256]int{}
	for i := range jewels {
		set[jewels[i]]++
	}

	ans := 0
	for i := range stones {
		if set[stones[i]] != 0 {
			ans++
		}
	}

	return ans
}

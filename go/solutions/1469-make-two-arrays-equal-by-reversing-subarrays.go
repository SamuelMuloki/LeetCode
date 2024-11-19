package solutions

func CanBeEqual2(target []int, arr []int) bool {
	count := make(map[int]int)
	for i := range target {
		count[target[i]]++
		count[arr[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

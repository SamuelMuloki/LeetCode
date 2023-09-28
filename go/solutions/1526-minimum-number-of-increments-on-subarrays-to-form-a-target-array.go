package solutions

func MinNumberOperations(target []int) int {
	st, count := []int{0}, 0
	for i := 0; i < len(target); i++ {
		if target[i] > st[len(st)-1] {
			count += target[i] - st[len(st)-1]
		}

		for st[len(st)-1] > target[i] {
			st = st[:len(st)-1]
		}
		st = append(st, target[i])
	}

	return count
}

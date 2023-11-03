package solutions

func BuildArray(target []int, n int) []string {
	st := make([]int, 0)
	ops := make([]string, 0)
	j := 0
	for i := 1; i <= n; i++ {
		if len(st) > 0 {
			if st[len(st)-1] != target[j] {
				st = st[:len(st)-1]
				ops = append(ops, "Pop")
			} else {
				j++
			}
		}

		if j >= len(target) {
			break
		}

		st = append(st, i)
		ops = append(ops, "Push")
	}

	return ops
}

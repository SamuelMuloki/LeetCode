package solutions

import "strconv"

func CalPoints(operations []string) int {
	st := []int{}
	for _, op := range operations {
		if op == "+" {
			last, second := st[len(st)-1], st[len(st)-2]
			st = append(st, last+second)
		} else if op == "D" {
			last := st[len(st)-1]
			st = append(st, last*2)
		} else if op == "C" {
			st = st[:len(st)-1]
		} else {
			num, _ := strconv.Atoi(op)
			st = append(st, num)
		}
	}

	ans := 0
	for i := range st {
		ans += st[i]
	}

	return ans
}

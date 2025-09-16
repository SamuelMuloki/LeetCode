package solutions

func ReplaceNonCoprimes(nums []int) []int {
	var gcd func(x, y int) int
	gcd = func(x, y int) int {
		if y == 0 {
			return x
		} else {
			return gcd(y, x%y)
		}
	}

	st := []int{}
	for _, num := range nums {
		st = append(st, num)
		for len(st) > 1 {
			a, b := st[len(st)-2], st[len(st)-1]
			g := gcd(a, b)
			if g == 1 {
				break
			}
			st = st[:len(st)-2]
			st = append(st, a*b/g)
		}
	}

	return st
}

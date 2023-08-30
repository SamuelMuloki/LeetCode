package solutions

func MySqrt(x int) int {
	l, r := 0, x

	for {
		mid := (l + r) / 2
		midS := mid * mid

		if midS <= x && x < (mid+1)*(mid+1) {
			return mid
		} else if midS < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}

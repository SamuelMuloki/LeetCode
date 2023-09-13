package solutions

func IsPerfectSquare(num int) bool {
	l, r := 0, num

	for l <= r {
		mid := l + (r-l)/2
		midS := mid * mid

		if midS == num {
			return true
		} else if midS < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

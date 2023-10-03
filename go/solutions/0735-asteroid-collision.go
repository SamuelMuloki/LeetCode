package solutions

func AsteroidCollision(asteroids []int) []int {
	st := []int{}
	var abs = func(ast int) int {
		if ast < 0 {
			return -ast
		}

		return ast
	}

	var sameSign = func(x, y int) bool {
		return x < 0 && y < 0 || x > 0 && y > 0
	}

	for i := 0; i < len(asteroids); i++ {
		for len(st) > 0 && !sameSign(asteroids[i], st[len(st)-1]) && asteroids[i] < 0 && abs(asteroids[i]) > abs(st[len(st)-1]) {
			st = st[:len(st)-1]
		}

		if len(st) > 0 && !sameSign(asteroids[i], st[len(st)-1]) {
			if abs(asteroids[i]) == st[len(st)-1] {
				st = st[:len(st)-1]
				continue
			}

			if abs(asteroids[i]) < st[len(st)-1] {
				continue
			}
		}

		st = append(st, asteroids[i])
	}

	return st
}

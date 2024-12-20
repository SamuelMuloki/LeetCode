package solutions

import "math"

func FindRotateSteps(ring string, key string) int {
	var m, n = len(ring), len(key)
	var pos = make(map[uint8][]int)
	for i := 0; i < len(ring); i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}
	var mem = make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if j == n {
			return
		}
		var p = &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = math.MaxInt
		for _, po := range pos[key[j]] {
			res = min(res, dfs(po, j+1)+min(abs(i-po), m-abs(i-po)))
		}
		return
	}
	return n + dfs(0, 0)
}

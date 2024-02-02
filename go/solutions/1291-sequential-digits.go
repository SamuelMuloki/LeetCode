package solutions

import "sort"

func SequentialDigits(low int, high int) []int {
	ans := make([]int, 0)
	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			if num >= low && num <= high {
				ans = append(ans, num)
			}
		}
	}

	sort.Ints(ans)
	return ans
}

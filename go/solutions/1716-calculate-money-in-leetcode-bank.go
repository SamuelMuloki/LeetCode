package solutions

func TotalMoney(n int) int {
	ans, monday := 0, 1
	for num := n; num > 0; num -= 7 {
		for day := 0; day < min(num, 7); day++ {
			ans += monday + day
		}
		monday++
	}
	return ans
}

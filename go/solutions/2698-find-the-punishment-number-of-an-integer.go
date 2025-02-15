package solutions

func PunishmentNumber(n int) int {
	ans := 0
	for num := 1; num <= n; num++ {
		mul := num * num
		if canPartition(mul, num) {
			ans += mul
		}
	}

	return ans
}

func canPartition(mul, target int) bool {
	if target < 0 || mul < target {
		return false
	}

	if mul == target {
		return true
	}

	return canPartition(mul/10, target-(mul%10)) ||
		canPartition(mul/100, target-(mul%100)) ||
		canPartition(mul/1000, target-(mul%1000))
}

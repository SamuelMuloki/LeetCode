package solutions

import "strconv"

func PunishmentNumber(n int) int {
	ans := 0
	for num := 1; num <= n; num++ {
		mul := num * num
		strNum := strconv.Itoa(mul)
		if canPartition(strNum, num) {
			ans += mul
		}
	}

	return ans
}

func canPartition(strNum string, target int) bool {
	if strNum == "" && target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	for idx := 0; idx < len(strNum); idx++ {
		left := strNum[0 : idx+1]
		right := strNum[idx+1 : len(strNum)]
		leftNum, _ := strconv.Atoi(left)

		if canPartition(right, target-leftNum) {
			return true
		}
	}

	return false
}

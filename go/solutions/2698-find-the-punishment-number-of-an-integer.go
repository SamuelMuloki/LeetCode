package solutions

import "strconv"

func PunishmentNumber(n int) int {
	ans := 0
	for num := 1; num <= n; num++ {
		mul := num * num
		strNum := strconv.Itoa(mul)
		memo := make([][]int, len(strNum))
		for i := range memo {
			memo[i] = make([]int, num+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		if findPartitions(0, 0, strNum, num, memo) {
			ans += mul
		}
	}

	return ans
}

func findPartitions(start, sum int, strNum string, target int, memo [][]int) bool {
	if start == len(strNum) {
		return sum == target
	}

	if sum > target {
		return false
	}

	if memo[start][sum] != -1 {
		return memo[start][sum] == 1
	}

	partitionFound := false
	for curr := start; curr < len(strNum); curr++ {
		currStr := strNum[start : curr+1]
		addEnd, _ := strconv.Atoi(currStr)

		partitionFound = partitionFound || findPartitions(curr+1, sum+addEnd, strNum, target, memo)
		if partitionFound {
			memo[start][sum] = 1
			return true
		}
	}

	memo[start][sum] = 0
	return false
}

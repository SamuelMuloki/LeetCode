package solutions

func ConstructDistancedSequence(n int) []int {
	var findSequence func(currIndex int, res []int, isNumberUsed []bool) bool
	findSequence = func(currIndex int, res []int, isNumberUsed []bool) bool {
		if currIndex == len(res) {
			return true
		}

		if res[currIndex] != 0 {
			return findSequence(currIndex+1, res, isNumberUsed)
		}

		for num := n; num >= 1; num-- {
			if isNumberUsed[num] {
				continue
			}

			isNumberUsed[num] = true
			res[currIndex] = num

			if num == 1 {
				if findSequence(currIndex+1, res, isNumberUsed) {
					return true
				}
			} else if currIndex+num < len(res) && res[currIndex+num] == 0 {
				res[currIndex+num] = num
				if findSequence(currIndex+1, res, isNumberUsed) {
					return true
				}

				res[currIndex+num] = 0
			}

			isNumberUsed[num] = false
			res[currIndex] = 0
		}

		return false
	}

	ans := make([]int, (2*n)-1)
	isNumberUsed := make([]bool, n+1)
	findSequence(0, ans, isNumberUsed)

	return ans
}

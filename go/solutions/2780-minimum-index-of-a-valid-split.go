package solutions

func MinimumIndex(nums []int) int {
	n := len(nums)
	set := make(map[int]int)
	majorityNum, maxCnt := 0, 0
	for _, num := range nums {
		set[num]++
		if set[num] > set[majorityNum] {
			majorityNum = num
		}

		maxCnt = max(maxCnt, set[num])
	}

	var check = func(i, left, right int) bool {
		return left*2 > i+1 && right*2 > n-i-1
	}

	left, right := 0, maxCnt
	for i := 0; i < len(nums); i++ {
		if nums[i] == majorityNum {
			left++
			right--
		}

		if check(i, left, right) {
			return i
		}
	}

	return -1
}

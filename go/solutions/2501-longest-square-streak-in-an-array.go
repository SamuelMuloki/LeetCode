package solutions

func LongestSquareStreak(nums []int) int {
	arr := [100001]int{}
	for i := range nums {
		arr[nums[i]] = nums[i]
	}

	ans := 0
	for i := range nums {
		count := 1
		for num := nums[i] * nums[i]; num < len(arr); num = num * num {
			if arr[num] == 0 {
				break
			}
			count++
		}
		ans = max(ans, count)
	}

	if ans == 1 {
		return -1
	}

	return ans
}

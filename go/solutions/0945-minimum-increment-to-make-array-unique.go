package solutions

func MinIncrementForUnique(nums []int) int {
	ans := 0
	count := [200000]int{}
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		for count[curr] == 1 {
			curr++
		}
		count[curr]++
		ans += curr - nums[i]
	}

	return ans
}

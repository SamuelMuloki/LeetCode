package leetcode

func ContainsDuplicate(nums []int) bool {
	times := make(map[int]int)

	for _, num := range nums {
		if times[num] == 1 {
			return true
		}
		times[num] = times[num] + 1
	}
	return false
}

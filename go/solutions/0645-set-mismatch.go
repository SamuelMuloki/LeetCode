package solutions

func FindErrorNums(nums []int) []int {
	n := len(nums)
	arr := make([]int, n+1)
	var rep int
	for i := range nums {
		arr[nums[i]]++
		if arr[nums[i]] == 2 {
			rep = nums[i]
		}
	}

	for i := 1; i <= n; i++ {
		if arr[i] == 0 {
			return []int{rep, i}
		}
	}

	return nil
}

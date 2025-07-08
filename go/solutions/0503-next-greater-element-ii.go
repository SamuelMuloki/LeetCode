package solutions

func NextGreaterElementsII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	st := []int{}
	for i := 2*n - 1; i > 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[i%n] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = nums[st[len(st)-1]]
		}
		st = append(st, i%n)
	}

	return res
}

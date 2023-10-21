package solutions

func MaxSlidingWindow(nums []int, k int) []int {
	var queue []int
	ans := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}

		if queue[0] == i-k+1 {
			queue = queue[1:]
		}
	}

	return ans
}

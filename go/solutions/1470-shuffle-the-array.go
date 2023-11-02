package solutions

func Shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	k, j := n, 0
	for i := 0; i < len(nums); i++ {
		if i&1 == 0 {
			ans[i] = nums[j]
			j++
		} else {
			ans[i] = nums[k]
			k++
		}
	}

	return ans
}

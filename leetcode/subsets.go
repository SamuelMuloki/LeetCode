package leetcode

func Subsets(nums []int) [][]int {
	subs := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		subs = append(subs, append([]int{}, curr...))
		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return subs
}

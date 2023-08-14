package leetcode

func Subsets(nums []int) [][]int {
	subs := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		addToBase := [][]int{}
		for j := 0; j < len(subs); j++ {
			newSet := append([]int{}, subs[j]...)
			newSet = append(newSet, nums[i])
			addToBase = append(addToBase, newSet)
		}

		subs = append(subs, addToBase...)
	}

	return subs
}

package solutions

func Search(nums []int, target int) int {
	index := -1

	mid := len(nums) / 2
	switch {
	case len(nums) == 0:
		return index
	case nums[mid] > target:
		index = Search(nums[:mid], target)
	case nums[mid] < target:
		index = Search(nums[mid+1:], target)
		if index >= 0 {
			index += mid + 1
		}
	default:
		index = mid
	}

	return index
}

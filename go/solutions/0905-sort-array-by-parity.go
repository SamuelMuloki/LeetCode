package solutions

func SortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		isIEven, isJEven := nums[i]%2 == 0, nums[j]%2 == 0
		if !isIEven && isJEven {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j-1
		} else if isIEven {
			i++
		} else if !isJEven {
			j--
		}
	}

	return nums
}

package solutions

func MaxProduct(nums []int) int {
	maxProd := nums[0]
	for i, minVal, maxVal := 1, maxProd, maxProd; i < len(nums); i++ {
		if nums[i] < 0 {
			minVal, maxVal = maxVal, minVal
		}

		minVal = min(nums[i], minVal*nums[i])
		maxVal = max(nums[i], maxVal*nums[i])
		maxProd = max(maxProd, maxVal)
	}

	return maxProd
}

package solutions

func FindDifferentBinaryString(nums []string) string {
	ans := ""
	for i := range nums {
		val, curr := "0", nums[i][i]
		if curr == '0' {
			val = "1"
		}
		ans += val
	}

	return ans
}

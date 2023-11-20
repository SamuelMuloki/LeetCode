package solutions

func CountDistinctIntegers(nums []int) int {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		reversed, temp := 0, nums[i]
		for temp > 0 {
			reversed = (reversed * 10) + (temp % 10)
			temp /= 10
		}

		set[nums[i]]++
		set[reversed]++
	}

	return len(set)
}

package solutions

func MissingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res = res ^ i ^ nums[i] // a ^ b ^ b = a
	}

	return res
}

package solutions

func SumIndicesWithKSetBits(nums []int, k int) int {
	bin := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		bin[i] = bin[i>>1] + i&1
	}

	sum := 0
	for i := 0; i < len(bin); i++ {
		if bin[i] == k {
			sum += nums[i]
		}
	}

	return sum
}

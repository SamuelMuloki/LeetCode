package solutions

func NumOfSubarrays(arr []int) int {
	sum := 0
	count := [2]int{1, 0}
	for i := 0; i < len(arr); i++ {
		sum = ((sum + arr[i]) % 2) % 2
		count[sum]++
	}

	return (count[0] * count[1]) % (1e9 + 7)
}

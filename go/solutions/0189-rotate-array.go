package solutions

func RotateArray(nums []int, k int) {
	n := len(nums)
	arr := append([]int{}, nums...)
	for i := 0; i < n; i++ {
		nums[(i+k)%n] = arr[i]
	}
}

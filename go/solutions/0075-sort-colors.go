package solutions

func SortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			swapArr(nums, i, j)
		}
	}

	swapArr(nums, i+1, high)

	return i + 1
}

func swapArr(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

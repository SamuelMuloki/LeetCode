package solutions

func IsArraySpecial2(nums []int, queries [][]int) []bool {
	arr := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			arr = append(arr, i)
		}
	}

	ans := make([]bool, len(queries))
	for j, query := range queries {
		start := query[0]
		end := query[1]
		foundIdx := binarySearch2(start+1, end, arr)
		if foundIdx {
			ans[j] = false
		} else {
			ans[j] = true
		}
	}

	return ans
}

func binarySearch2(start, end int, arr []int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) >> 1
		idx := arr[mid]
		if idx < start {
			left = mid + 1
		} else if idx > end {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

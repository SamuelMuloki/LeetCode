package solutions

func FindSmallestInteger(nums []int, value int) int {
	modCount := make([]int, value)
	for _, num := range nums {
		r := (num%value + value) % value
		modCount[r]++
	}

	for i := 0; ; i++ {
		if modCount[i%value] == 0 {
			return i
		}
		modCount[i%value]--
	}
}

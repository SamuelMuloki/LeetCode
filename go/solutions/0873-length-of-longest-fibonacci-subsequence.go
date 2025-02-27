package solutions

func LenLongestFibSubseq(arr []int) int {
	n := len(arr)
	set := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			prev, curr := arr[j], arr[i]+arr[j]
			cnt := 2

			for set[curr] {
				tmp := curr
				curr += prev
				prev = tmp
				cnt++
				ans = max(ans, cnt)
			}
		}
	}

	return ans
}

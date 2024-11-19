package solutions

func NumTeams(rating []int) int {
	n := len(rating)
	teams := 0

	for mid := 0; mid < n; mid++ {
		leftSmaller := 0
		rightLarger := 0

		for left := mid - 1; left >= 0; left-- {
			if rating[left] < rating[mid] {
				leftSmaller++
			}
		}

		for right := mid + 1; right < n; right++ {
			if rating[right] > rating[mid] {
				rightLarger++
			}
		}

		teams += leftSmaller * rightLarger

		leftLarger := mid - leftSmaller
		rightSmaller := n - mid - 1 - rightLarger

		teams += leftLarger * rightSmaller
	}

	return teams
}

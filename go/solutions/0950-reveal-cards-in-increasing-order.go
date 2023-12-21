package solutions

import "sort"

func DeckRevealedIncreasing(deck []int) []int {
	queue := append([]int{}, deck...)
	sort.Ints(queue)
	ans := make([]int, 0)
	for len(queue) > 0 {
		ans = append([]int{queue[len(queue)-1]}, ans...)
		queue = queue[:len(queue)-1]
		ans = append([]int{ans[len(ans)-1]}, ans...)
		ans = ans[:len(ans)-1]
	}
	ans = append(ans, ans[0])
	ans = ans[1:]

	return ans
}

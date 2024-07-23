package solutions

func Trap(height []int) int {
	ans, n := 0, len(height)
	maxLeft, maxRight := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for j := n - 2; j >= 0; j-- {
		maxRight[j] = max(maxRight[j+1], height[j+1])
	}
	for i := 1; i < n-1; i++ {
		if height[i] < maxLeft[i] && height[i] < maxRight[i] {
			ans += min(maxLeft[i], maxRight[i]) - height[i]
		}
	}
	return ans
}

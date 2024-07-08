package solutions

func PassThePillow(n int, time int) int {
	chunks := time / (n - 1)
	if chunks%2 == 0 {
		return (time % (n - 1)) + 1
	} else {
		return n - (time % (n - 1))
	}
}

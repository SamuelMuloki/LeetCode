package solutions

func BitwiseComplement(n int) int {
	x := 1
	for ; x < n; x = (x << 1) + 1 {
	}
	return n ^ x
}

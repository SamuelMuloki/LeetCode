package solutions

func FindComplement(num int) int {
	x := 1
	for ; x < num; x = (x << 1) + 1 {
	}
	return num ^ x
}

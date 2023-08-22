// https://solutions.com/problems/powx-n
package solutions

func MyPow(x float64, n int) float64 {
	return binaryExp(x, n)
}

func binaryExp(x float64, n int) float64 {
	if n == 0 {
		return float64(1)
	}

	if n < 0 {
		return float64(1) / binaryExp(x, -n)
	}

	if n%2 != 0 {
		return x * binaryExp(x*x, (n-1)/2)
	}

	return binaryExp(x*x, n/2)
}

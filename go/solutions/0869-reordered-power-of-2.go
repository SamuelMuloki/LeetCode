package solutions

import "strconv"

func ReorderedPowerOf2(n int) bool {
	var isPowerOfTwo = func(num int) bool {
		return num > 0 && (num-1)&num == 0
	}

	res := false
	var permute func(arr []rune, l, r int) bool
	permute = func(arr []rune, l, r int) bool {
		if l == r {
			num, _ := strconv.Atoi(string(arr))
			res = res || arr[0] != '0' && isPowerOfTwo(num)
			return res
		}

		for i := l; i <= r; i++ {
			arr[l], arr[i] = arr[i], arr[l]
			permute(arr, l+1, r)
			arr[l], arr[i] = arr[i], arr[l]
		}

		return res
	}

	str := strconv.Itoa(n)
	runes := []rune(str)

	return permute(runes, 0, len(runes)-1)
}

// https://leetcode.com/problems/minimum-operations-to-make-a-special-number
package solutions

func MinimumOperations(num string) int {
	fiveFound, zeroFound := false, false
	for i := len(num) - 1; i >= 0; i-- {
		switch {
		case zeroFound && num[i] == '0':
			return len(num) - 2 - i
		case zeroFound && num[i] == '5':
			return len(num) - 2 - i
		case fiveFound && num[i] == '2':
			return len(num) - 2 - i
		case fiveFound && num[i] == '7':
			return len(num) - 2 - i
		case num[i] == '5':
			fiveFound = true
		case num[i] == '0':
			zeroFound = true
		}
	}

	if zeroFound {
		return len(num) - 1
	}

	return len(num)
}

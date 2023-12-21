package solutions

import "strconv"

func CountAndSay(n int) string {
	i, say := 0, "1"
	for ; n > 1; n-- {
		count, str := 1, ""
		for i = 1; i < len(say); i++ {
			if say[i-1] != say[i] {
				str += strconv.Itoa(count) + string(say[i-1])
				count = 1
			} else {
				count++
			}
		}

		str += strconv.Itoa(count) + string(say[i-1])
		say = str
	}

	return say
}

package solutions

func OpenLock(deadends []string, target string) int {
	invalid := make(map[string]bool)
	for _, deadend := range deadends {
		invalid[deadend] = true
	}

	visited := make(map[string]bool)
	queue := [][]string{{"0000"}}
	diff := []int{1, -1}

	if invalid["0000"] {
		return -1
	}

	if target == "0000" {
		return 0
	}

	steps := 0

	for len(queue) > 0 {
		codes := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		newCodes := make([]string, 0)

		for _, code := range codes {
			for i := 0; i < 4; i++ {
				for _, d := range diff {
					newchar := int(code[i]-'0') + d

					if newchar == 10 {
						newchar = 0
					} else if newchar == -1 {
						newchar = 9
					}

					chars := make([]byte, 4)
					copy(chars, code)
					chars[i] = byte(newchar) + '0'
					str := string(chars)

					if str == target {
						return steps + 1
					}

					if visited[str] || invalid[str] {
						continue
					}

					visited[str] = true
					newCodes = append(newCodes, str)
				}
			}
		}

		if len(newCodes) > 0 {
			queue = append(queue, newCodes)
		}

		steps++
	}

	return -1
}

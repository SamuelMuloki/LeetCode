package solutions

func MinOperations5(logs []string) int {
	st := []string{}
	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" {
			continue
		} else if logs[i] == "../" {
			if len(st) == 0 {
				continue
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, logs[i])
		}
	}

	return len(st)
}

package solutions

func NumUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, email := range emails {
		str := ""
		for i := 0; i < len(email); {
			if email[i] == '.' {
				i++
				continue
			}

			if email[i] == '+' {
				for ; email[i] != '@'; i++ {
				}
			}

			if email[i] == '@' {
				str += string(email[i:])
				break
			}

			str += string(email[i])
			i++
		}

		set[str] = true
	}

	return len(set)
}

package solutions

func ParseBoolExpr(expression string) bool {
	st := []rune{}
	for _, expr := range expression {
		if expr == '(' || expr == ',' {
			continue
		}

		if expr == '!' || expr == '&' || expr == '|' || expr == 'f' || expr == 't' {
			st = append(st, expr)
		} else if expr == ')' {
			hasTrue, hasFalse := false, false
			for st[len(st)-1] != '!' && st[len(st)-1] != '&' && st[len(st)-1] != '|' {
				topValue := st[len(st)-1]
				st = st[:len(st)-1]
				if topValue == 't' {
					hasTrue = true
				}
				if topValue == 'f' {
					hasFalse = true
				}
			}

			top := st[len(st)-1]
			st = st[:len(st)-1]
			if top == '!' {
				if hasTrue {
					st = append(st, 'f')
				} else {
					st = append(st, 't')
				}
			} else if top == '&' {
				if hasFalse {
					st = append(st, 'f')
				} else {
					st = append(st, 't')
				}
			} else {
				if hasTrue {
					st = append(st, 't')
				} else {
					st = append(st, 'f')
				}
			}
		}
	}

	return st[len(st)-1] == 't'
}

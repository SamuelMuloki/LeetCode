package solutions

import (
	"sort"
	"strconv"
)

func CountOfAtoms(formula string) string {
	formulas := []map[string]int{make(map[string]int)}
	var i int
	for i < len(formula) {
		c := formula[i]
		if c == '(' {
			formulas = append(formulas, make(map[string]int))
			i++
		} else if c == ')' {
			curMap := formulas[len(formulas)-1]
			formulas = formulas[:len(formulas)-1]
			i++

			var multiplier []byte
			for i < len(formula) && isDigit(rune(formula[i])) {
				multiplier = append(multiplier, formula[i])
				i++
			}

			if len(multiplier) != 0 {
				v, _ := strconv.Atoi(string(multiplier))
				for k := range curMap {
					curMap[k] *= v
				}
			}

			for k, v := range curMap {
				formulas[len(formulas)-1][k] += v
			}
		} else {
			curAtom := []byte{c}
			i++
			for i < len(formula) && formula[i] >= 'a' && formula[i] <= 'z' {
				curAtom = append(curAtom, formula[i])
				i++
			}

			var curDigit []byte
			for i < len(formula) && isDigit(rune(formula[i])) {
				curDigit = append(curDigit, formula[i])
				i++
			}

			if len(curDigit) == 0 {
				formulas[len(formulas)-1][string(curAtom)]++
			} else {
				v, _ := strconv.Atoi(string(curDigit))
				formulas[len(formulas)-1][string(curAtom)] += v
			}
		}
	}

	var keys []string
	for k := range formulas[len(formulas)-1] {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var a string
	for _, s := range keys {
		a += s
		if formulas[len(formulas)-1][s] > 1 {
			a += strconv.Itoa(formulas[len(formulas)-1][s])
		}
	}

	return a
}

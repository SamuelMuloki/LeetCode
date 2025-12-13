package solutions

import (
	"sort"
	"strings"
)

func ValidateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var isValid = func(ch rune) bool {
		return ch >= 'a' && ch <= 'z' ||
			ch >= 'A' && ch <= 'Z' ||
			ch >= '0' && ch <= '9' || ch == '_'
	}

	var isValidCode = func(word string) bool {
		if len(word) == 0 {
			return false
		}

		for _, ch := range word {
			if !isValid(ch) {
				return false
			}
		}

		return true
	}

	order := map[string]int{
		"electronics": 1,
		"grocery":     2,
		"pharmacy":    3,
		"restaurant":  4,
	}

	type Order struct {
		code string
		pos  int
	}

	arr := []Order{}
	for i, c := range code {
		var num int
		if v, ok := order[strings.ToLower(businessLine[i])]; !ok {
			continue
		} else {
			num = v
		}

		if isActive[i] && isValidCode(c) {
			arr = append(arr, Order{
				code: c,
				pos:  num,
			})
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].pos < arr[j].pos {
			return true
		}
		if arr[i].pos > arr[j].pos {
			return false
		}
		return arr[i].code < arr[j].code
	})

	res := []string{}
	for _, order := range arr {
		res = append(res, order.code)
	}

	return res
}

package utils

import "testing"

type testwords struct {
	words, result []string
}

func TestSwap(t *testing.T) {
	tests := []testwords{
		{[]string{"Hello", "world"}, []string{"world", "Hello"}},
		{[]string{"Hello", "世界"}, []string{"世界", "Hello"}},
	}
	for _, c := range tests {
		r1, r2 := Swap(c.words[0], c.words[1])

		if r1 != c.result[0] || r2 != c.result[1] {
			t.Error(
				"Expected",
				c.result[0],
				c.result[1],
				"got",
				r1,
				r2,
			)
		}
	}
}

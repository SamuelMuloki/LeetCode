package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec2 struct{}

func NewCodec2() *Codec2 {
	return &Codec2{}
}

func (this *Codec2) Encode(strs []string) string {
	sizes := []string{}
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return fmt.Sprintf("%s%s%s", strings.Join(sizes, ","), "#", strings.Join(strs, ""))
}

func (this *Codec2) Decode(s string) []string {
	res := []string{}
	if s == "" {
		return res
	}

	parts := strings.SplitN(s, "#", 2)
	sizes := strings.Split(parts[0], ",")
	idx := 0
	for i := range sizes {
		size, _ := strconv.Atoi(sizes[i])
		res = append(res, string(parts[1][idx:idx+size]))
		idx += size
	}

	return res
}

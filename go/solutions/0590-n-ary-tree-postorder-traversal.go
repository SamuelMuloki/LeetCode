package solutions

import (
	"container/list"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func Postorder(root *utils.Node) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}

	s := list.New()
	s.PushFront(root)
	for s.Len() > 0 {
		curr := s.Remove(s.Front()).(*utils.Node)
		if curr != nil {
			output = append(output, curr.Val)
		}

		for _, child := range curr.Children {
			s.PushFront(child)
		}
	}

	reverse(output)

	return output
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

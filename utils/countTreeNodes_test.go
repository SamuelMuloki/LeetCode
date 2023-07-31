package utils

import (
	"testing"

	"github.com/SamuelMuloki/GOExamples/leetcode"
)

type TestNode struct {
	node   *leetcode.TreeNode
	result int
}

func TestCountNodes(t *testing.T) {
	testNodes := []TestNode{
		{
			node:   nil,
			result: 0,
		},
		{
			node: &leetcode.TreeNode{
				Val:   1,
				Left:  &leetcode.TreeNode{Val: 2},
				Right: &leetcode.TreeNode{Val: 3},
			},
			result: 3,
		},
	}

	for _, testNode := range testNodes {
		res := CountTreeNodes(testNode.node)

		if res != testNode.result {
			t.Error(
				"Expected",
				testNode.result,
				"got",
				res,
			)
		}
	}
}

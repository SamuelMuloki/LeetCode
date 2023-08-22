package utils

import (
	"testing"
)

type TestNode struct {
	node   *TreeNode
	result int
}

func TestCountNodes(t *testing.T) {
	testNodes := []TestNode{
		{
			node:   nil,
			result: 0,
		},
		{
			node: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
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

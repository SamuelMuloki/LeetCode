package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

type Path2Sum struct {
	currSum int
	elems   []int
	node    *utils.TreeNode
}

func PathSum2(root *utils.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	s := make([]Path2Sum, 0)
	s = append(s, Path2Sum{node: root, currSum: targetSum})

	for len(s) > 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]

		if last.node.Left == nil && last.node.Right == nil && last.currSum == last.node.Val {
			curr := append(last.elems, last.node.Val)
			res = append(res, append([]int{}, curr...))
		}

		if last.node.Right != nil {
			s = append(s, Path2Sum{node: last.node.Right, elems: append(last.elems, last.node.Val), currSum: last.currSum - last.node.Val})
		}

		if last.node.Left != nil {
			s = append(s, Path2Sum{node: last.node.Left, elems: append(last.elems, last.node.Val), currSum: last.currSum - last.node.Val})
		}
	}

	return res
}

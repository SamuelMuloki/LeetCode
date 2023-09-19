package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func Postorder(root *utils.Node) []int {
	output := make([]int, 0)
	recur(root, &output)

	return output
}

func recur(root *utils.Node, output *[]int) []int {
	if root == nil {
		return *output
	}

	if root.Children != nil {
		for _, child := range *root.Children {
			recur(&child, output)
		}
	}

	*output = append(*output, root.Val)

	return *output
}

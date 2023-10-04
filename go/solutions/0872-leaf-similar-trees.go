package solutions

import (
	"reflect"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LeafSimilar(root1 *utils.TreeNode, root2 *utils.TreeNode) bool {
	arr := make([][]int, 2)

	for i, val := range []*utils.TreeNode{root1, root2} {
		st := append([]*utils.TreeNode{}, val)
		for len(st) > 0 {
			last := st[len(st)-1]
			st = st[:len(st)-1]

			if last.Left == nil && last.Right == nil {
				arr[i] = append(arr[i], last.Val)
			}

			if last.Right != nil {
				st = append(st, last.Right)
			}

			if last.Left != nil {
				st = append(st, last.Left)
			}
		}
	}

	return reflect.DeepEqual(arr[0], arr[1])
}

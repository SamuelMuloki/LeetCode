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
	arr1, arr2 := make([]int, 0), make([]int, 0)

	find(root1, &arr1)
	find(root2, &arr2)

	return reflect.DeepEqual(arr1, arr2)
}

func find(root *utils.TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
		return
	}

	find(root.Left, arr)
	find(root.Right, arr)
}

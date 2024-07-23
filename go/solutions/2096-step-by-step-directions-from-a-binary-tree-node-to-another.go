package solutions

import (
	"strings"

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
func GetDirections(root *utils.TreeNode, startValue int, destValue int) string {
	var dfs func(root *utils.TreeNode, arr []byte, target int) []byte
	dfs = func(root *utils.TreeNode, arr []byte, target int) []byte {
		if root.Val == target {
			return arr
		}
		if root.Left != nil {
			newArr := append(arr, 'L')
			if found := dfs(root.Left, newArr, target); found != nil {
				return found
			}
		}
		if root.Right != nil {
			newArr := append(arr, 'R')
			if found := dfs(root.Right, newArr, target); found != nil {
				return found
			}
		}
		return nil
	}

	arrToS := dfs(root, []byte{}, startValue)
	arrToD := dfs(root, []byte{}, destValue)

	i := 0
	for len(arrToS) > i && len(arrToD) > i && arrToS[i] == arrToD[i] {
		i++
	}

	pathUp := strings.Repeat("U", len(arrToS)-i)
	pathDown := string(arrToD[i:])

	return pathUp + pathDown
}

package solutions

import (
	"sort"

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
func MinimumOperations2(root *utils.TreeNode) int {
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		first := len(queue)
		arr := make([]int, 0)
		for i := 0; i < first; i++ {
			arr = append(arr, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		ans += stOperations(arr)
		queue = queue[first:]
	}

	return ans
}

func stOperations(arr []int) int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	ans := 0
	pos := make(map[int]int)
	for i := 0; i < n; i++ {
		pos[arr[i]] = i
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedArr[i] {
			ans++

			currPos := pos[sortedArr[i]]
			pos[arr[i]] = currPos
			arr[currPos] = arr[i]
		}
	}

	return ans
}

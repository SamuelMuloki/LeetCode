package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeQuery struct {
	maxHeightAfterRemoval [100001]int
	currentMaxHeight      int
}

func TreeQueries(root *utils.TreeNode, queries []int) []int {
	tQuery := &TreeQuery{}

	tQuery.traverseLeftToRight(root, 0)

	tQuery.currentMaxHeight = 0
	tQuery.traverseRightToLeft(root, 0)

	queryResults := make([]int, len(queries))
	for i, query := range queries {
		queryResults[i] = tQuery.maxHeightAfterRemoval[query]
	}

	return queryResults
}

func (t *TreeQuery) traverseLeftToRight(node *utils.TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	t.maxHeightAfterRemoval[node.Val] = t.currentMaxHeight
	t.currentMaxHeight = max(t.currentMaxHeight, currentHeight)

	t.traverseLeftToRight(node.Left, currentHeight+1)
	t.traverseLeftToRight(node.Right, currentHeight+1)
}

func (t *TreeQuery) traverseRightToLeft(node *utils.TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	t.maxHeightAfterRemoval[node.Val] = max(
		t.maxHeightAfterRemoval[node.Val],
		t.currentMaxHeight,
	)

	t.currentMaxHeight = max(currentHeight, t.currentMaxHeight)

	t.traverseRightToLeft(node.Right, currentHeight+1)
	t.traverseRightToLeft(node.Left, currentHeight+1)
}

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
func CreateBinaryTree(descriptions [][]int) *utils.TreeNode {
	nodeMap := make(map[int]*utils.TreeNode)
	children := make(map[int]bool)
	for _, description := range descriptions {
		parentValue, childValue := description[0], description[1]
		isLeft := description[2] == 1
		if _, exists := nodeMap[parentValue]; !exists {
			nodeMap[parentValue] = &utils.TreeNode{Val: parentValue}
		}

		if _, exists := nodeMap[childValue]; !exists {
			nodeMap[childValue] = &utils.TreeNode{Val: childValue}
		}

		if isLeft {
			nodeMap[parentValue].Left = nodeMap[childValue]
		} else {
			nodeMap[parentValue].Right = nodeMap[childValue]
		}

		children[childValue] = true
	}

	for val, node := range nodeMap {
		if _, exists := children[val]; !exists {
			return node
		}
	}

	return nil
}

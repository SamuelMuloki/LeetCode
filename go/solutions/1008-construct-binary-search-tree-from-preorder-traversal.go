package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func BstFromPreorder(preorder []int) *utils.TreeNode {
	root := &utils.TreeNode{Val: preorder[0]}
	for i := range preorder {
		child := &utils.TreeNode{Val: preorder[i]}
		child.InsertNodeToTree(root)
	}

	return root
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func IsSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

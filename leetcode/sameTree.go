package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

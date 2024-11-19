package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func IsSubPath(head *utils.ListNode, root *utils.TreeNode) bool {
	var dfs func(head, cur *utils.ListNode, root *utils.TreeNode) bool
	dfs = func(head *utils.ListNode, cur *utils.ListNode, root *utils.TreeNode) bool {
		if cur == nil {
			return true
		}

		if root == nil {
			return false
		}

		if cur.Val == root.Val {
			cur = cur.Next
		} else if head.Val == root.Val {
			head = head.Next
		} else {
			cur = head
		}

		return dfs(head, cur, root.Left) || dfs(head, cur, root.Right)
	}

	return dfs(head, head, root)
}

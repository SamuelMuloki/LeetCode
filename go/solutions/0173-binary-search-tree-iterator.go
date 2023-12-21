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
type BSTIterator struct {
	st []int
}

func BSTIteratorConstructor(root *utils.TreeNode) BSTIterator {
	arr := make([]int, 0)
	s := make([]*utils.TreeNode, 0)
	curr := root

	for curr != nil || len(s) > 0 {
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}

		s, curr = s[:len(s)-1], s[len(s)-1]
		arr = append(arr, curr.Val)
		curr = curr.Right
	}

	return BSTIterator{arr}
}

func (this *BSTIterator) Next() int {
	val := this.st[0]
	this.st = this.st[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.st) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

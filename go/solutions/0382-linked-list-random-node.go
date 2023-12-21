package solutions

import (
	"math/rand"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type RandomNodeSolution struct {
	head *utils.ListNode
}

func RandomNodeConstructor(head *utils.ListNode) RandomNodeSolution {
	return RandomNodeSolution{head: head}
}

func (this *RandomNodeSolution) GetRandom() int {
	ans, n := 0, 0
	for curr := this.head; curr != nil; curr = curr.Next {
		n++
		if rand.Intn(n) == 0 {
			ans = curr.Val
		}
	}

	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

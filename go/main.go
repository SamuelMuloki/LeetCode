package main

import (
	"fmt"
	"strconv"

	"github.com/SamuelMuloki/LeetCode/go/methods"
	"github.com/SamuelMuloki/LeetCode/go/solutions"
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func helloWorld() {
	fmt.Println("hello world")
	fmt.Println("Hello, 世界")
}

func swap() {
	a, b := utils.Swap("hello", "world")
	fmt.Println(a, b)
}

func speak() {
	animals := []methods.Animal{methods.Dog{}, methods.Cat{}, methods.Llama{}, methods.JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	helloWorld()
	swap()
	speak()

	// Stack
	athletes := methods.New()
	athletes.Push(methods.Person{Name: "Lebron", Age: 37})
	athletes.Push(methods.Person{Name: "Jordan", Age: 50})
	athletes.Push(methods.Person{Name: "Kareem", Age: 65})
	athletes.Pop()
	fmt.Printf("There are %d GOATS\n", athletes.Length())

	// Linked List
	list := methods.LinkedList{}
	list.PushBack(&methods.Node{Data: 20})
	list.PushBack(&methods.Node{Data: 30})
	list.PushBack(&methods.Node{Data: 40})
	list.PushBack(&methods.Node{Data: 50})
	list.PushBack(&methods.Node{Data: 70})

	fmt.Printf("Length = %d\n", list.Length)
	list.Display()
	fmt.Printf("Length = %d\n", list.Length)
	list.Reverse()
	list.Display()

	// Queue
	queue := methods.Queue{Size: 3}
	fmt.Println(queue.Elements)
	queue.Enqueue(1)
	fmt.Println(queue.Elements)
	queue.Enqueue(2)
	fmt.Println(queue.Elements)
	queue.Enqueue(3)
	fmt.Println(queue.Elements)
	queue.Enqueue(5)
	fmt.Println(queue.Elements)
	elem := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue.Elements)
	queue.Enqueue(9)
	fmt.Println(queue.Elements)
	elem = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue.Elements)

	nums := []int{1, 2, 3, 1}
	fmt.Printf("Array contains duplicate: %v\n", solutions.ContainsDuplicate(nums))
	fmt.Printf("Is Anagram: %v\n", solutions.IsAnagram("anagram", "nagaram"))
	fmt.Printf("Two sum indices are %v\n", solutions.TwoSum([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("Group Anagrams: %v\n", solutions.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("The top k frequent elements are: %v\n", solutions.TopKFrequent([]int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}, 4))
	fmt.Printf("The product of array except self is: %v\n", solutions.ProductExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("The string is a valid palidronme: %v\n", solutions.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("Two sum 2 indices are %v\n", solutions.TwoSum2([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("The distinct triplets are %v\n", solutions.ThreeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Printf("The maximum area of the container is %v\n", solutions.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3}))
	fmt.Printf("The maximum sum of the array is %d\n", utils.MaxSum([]int{5, 2, -1, 0, 3}, 3))
	fmt.Printf("The maximum profit is %d\n", solutions.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("The length of the longest substring is %d\n", solutions.LengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("String has valid parentheses: %v\n", solutions.IsValid("()[]{}"))
	fmt.Printf("The index of the target is: %v\n", solutions.Search([]int{-1, 0, 3, 5, 9, 12}, 9))

	head := &solutions.ListNode{Val: 1, Next: &solutions.ListNode{Val: 2, Next: nil}}
	fmt.Printf("The Reversed linked list is: %v\n", solutions.ReverseList(head))

	list1 := &solutions.ListNode{Val: 1, Next: &solutions.ListNode{Val: 2, Next: &solutions.ListNode{Val: 4, Next: nil}}}
	list2 := &solutions.ListNode{Val: 1, Next: &solutions.ListNode{Val: 3, Next: &solutions.ListNode{Val: 4, Next: nil}}}
	fmt.Printf("The final lists after merging two lists is %v\n", solutions.MergeTwoLists(list1, list2))
	fmt.Printf("You can climb to the top in %d distinct ways\n", solutions.ClimbStairs(6))
	fmt.Printf("The median of two sorted arrays a, b is %v\n", solutions.FindMedianSortedArrays([]int{1, 2}, []int{3}))

	leftTree := &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}}}
	rightTree := &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}}}
	fmt.Printf("The trees p and q are equal: %v\n", solutions.IsSameTree(leftTree, rightTree))

	symTree := &utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}, Left: &utils.TreeNode{Val: 4}},
		Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 4}, Left: &utils.TreeNode{Val: 3}},
	}
	fmt.Printf("The root if binary tree is symetric: %v\n", solutions.IsSymmetric(symTree))

	stream := solutions.Stream{
		Set:      map[int]int{},
		MaxValue: 1,
	}
	stream.GetHighestNumber(1)
	stream.GetHighestNumber(2)
	stream.GetHighestNumber(3)
	fmt.Printf("The highest number for stream is %d\n", stream.GetHighestNumber(7))

	reorderHead := &solutions.ListNode{
		Val: 1, Next: &solutions.ListNode{Val: 2, Next: &solutions.ListNode{
			Val: 3, Next: &solutions.ListNode{Val: 4, Next: nil},
		}}}

	solutions.ReorderList(reorderHead)
	fmt.Printf("After reordering list head is now %v\n", reorderHead)
	fmt.Printf("The inverse of tree is :%v\n", solutions.InvertTree(symTree))

	depthTree := &utils.TreeNode{
		Val: 1,
		Left: &utils.TreeNode{
			Val: 2,
		},
	}
	fmt.Printf("The maximum depth of the binary tree is: %d\n", solutions.MaxDepth(depthTree))
	fmt.Printf("The diameter of the binary tree is: %d\n", solutions.DiameterOfBinaryTree(depthTree))
	fmt.Printf("Is the binary tree balanced: %v\n", solutions.IsBalanced(depthTree))

	kthLargest := solutions.Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf("The kth largest in the stream is: %v\n", kthLargest.Add(3))
	fmt.Printf("The last stone weight is: %v\n", solutions.LastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

	num := 19
	fmt.Printf("Number %d is a happy number: %v\n", num, solutions.IsHappy(num))

	digits := []int{9}
	fmt.Printf("Adding plus one equals: %v\n", solutions.PlusOne(digits))

	fmt.Printf("The value of myPow is %v\n", solutions.MyPow(2.00000, -2))
	fmt.Printf("The result of adding two string numbers is %s\n", solutions.AddStrings("289", "99"))
	fmt.Printf("The product of two string numbers is %s\n", solutions.Multiply("3866762897776739956", "15975363164662"))
	fmt.Printf("The single number in the array is %d\n", solutions.SingleNumber([]int{2, 2, 1}))
	fmt.Printf("The number is a power of two: %v\n", solutions.IsPowerOfTwo(0))

	numWeight, _ := strconv.ParseInt("00000000000000000000000010000000", 2, 64)
	fmt.Printf("The hamming weight of the number is %d\n", solutions.HammingWeight(uint32(numWeight)))

	fmt.Printf("The subsets of the array are %v\n", solutions.Subsets([]int{1, 2, 3}))
	fmt.Printf("The combination sum is %v\n", solutions.CombinationSum([]int{2, 3, 5}, 8))
	fmt.Printf("The permutation is %v\n", solutions.Permute([]int{1, 2, 3}))
	fmt.Printf("The subsets with dupes are %v\n", solutions.SubsetsWithDup([]int{4, 4, 4, 1, 4}))

	fmt.Printf("The combination sum 2 is %v\n", solutions.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

	var board [][]byte = [][]byte{
		{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')},
	}
	word := "ABCCED"
	fmt.Printf("Word exists in grid: %v\n", solutions.Exist(board, word))

	rootTree := utils.NewTree([]int{4, 3, 5, 1, 2})
	subTree := utils.NewTree([]int{3, 1, 2})

	fmt.Printf("There is a subtree of the root tree: %v\n", solutions.IsSubtree(rootTree, subTree))
	fmt.Printf("The possible palindrome partitions of s are %v\n", solutions.Partition("aab"))
	fmt.Printf("The letter combnations are %v\n", solutions.LetterCombinations("23"))
	fmt.Printf("Converting columnNumber to title equals: %s\n", solutions.ConvertToTitle(2147483647))

	arrBytes := []byte{'h', 'e', 'l', 'l', 'o'}
	solutions.ReverseString(arrBytes)
	fmt.Printf("The result of reversing the array is %v\n", arrBytes)
	fmt.Printf("The distinct solutions for the n-queens puzzle are, %v\n", solutions.SolveNQueens(4))
	fmt.Printf("The minimum element in the rotated array is %d\n", solutions.FindMin([]int{4, 5, 1, 2, 3}))
	fmt.Printf("The index of the element in the rotated array is %d\n", solutions.RotatedSearch([]int{4, 5, 6, 7, 0, 1, 2}, 6))
}

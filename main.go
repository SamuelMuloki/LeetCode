package main

import (
	"fmt"

	"github.com/SamuelMuloki/GOExamples/leetcode"
	"github.com/SamuelMuloki/GOExamples/methods"
	"github.com/SamuelMuloki/GOExamples/utils"
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
	fmt.Printf("Array contains duplicate: %v\n", leetcode.ContainsDuplicate(nums))
	fmt.Printf("Is Anagram: %v\n", leetcode.IsAnagram("anagram", "nagaram"))
	fmt.Printf("Two sum indices are %v\n", leetcode.TwoSum([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("Group Anagrams: %v\n", leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("The top k frequent elements are: %v\n", leetcode.TopKFrequent([]int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}, 4))
	fmt.Printf("The product of array except self is: %v\n", leetcode.ProductExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("The string is a valid palidronme: %v\n", leetcode.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("Two sum 2 indices are %v\n", leetcode.TwoSum2([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("The distinct triplets are %v\n", leetcode.ThreeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Printf("The maximum area of the container is %v\n", leetcode.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3}))
	fmt.Printf("The maximum sum of the array is %d\n", leetcode.MaxSum([]int{5, 2, -1, 0, 3}, 3))
	fmt.Printf("The maximum profit is %d\n", leetcode.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("The length of the longest substring is %d\n", leetcode.LengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("String has valid parentheses: %v\n", leetcode.IsValid("()[]{}"))
	fmt.Printf("The index of the target is: %v\n", leetcode.Search([]int{-1, 0, 3, 5, 9, 12}, 9))

	head := &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2, Next: nil}}
	fmt.Printf("The Reversed linked list is: %v\n", leetcode.ReverseList(head))

	list1 := &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 4, Next: nil}}}
	list2 := &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 3, Next: &leetcode.ListNode{Val: 4, Next: nil}}}
	fmt.Printf("The final lists after merging two lists is %v\n", leetcode.MergeTwoLists(list1, list2))
	fmt.Printf("You can climb to the top in %d distinct ways\n", leetcode.ClimbStairs(6))
	fmt.Printf("The median of two sorted arrays a, b is %v\n", leetcode.FindMedianSortedArrays([]int{1, 2}, []int{3}))
}

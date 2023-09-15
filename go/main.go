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

	// DFS Graph
	g := methods.NewDFSGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	g.DFS(2)

	nums := []int{1, 2, 3, 1}
	fmt.Printf("Array contains duplicate: %t\n", solutions.ContainsDuplicate(nums))
	fmt.Printf("Is Anagram: %t\n", solutions.IsAnagram("anagram", "nagaram"))
	fmt.Printf("Two sum indices are %v\n", solutions.TwoSum([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("Group Anagrams: %v\n", solutions.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("The top k frequent elements are: %v\n", solutions.TopKFrequent([]int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}, 4))
	fmt.Printf("The product of array except self is: %v\n", solutions.ProductExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("The string is a valid palidronme: %t\n", solutions.IsPalindrome("A man, a plan, a canal: Panama"))
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
	fmt.Printf("The trees p and q are equal: %t\n", solutions.IsSameTree(leftTree, rightTree))

	symTree := &utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}, Left: &utils.TreeNode{Val: 4}},
		Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 4}, Left: &utils.TreeNode{Val: 3}},
	}
	fmt.Printf("The root if binary tree is symetric: %t\n", solutions.IsSymmetric(symTree))

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
	fmt.Printf("Is the binary tree balanced: %t\n", solutions.IsBalanced(depthTree))

	kthLargest := solutions.Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf("The kth largest in the stream is: %v\n", kthLargest.Add(3))
	fmt.Printf("The last stone weight is: %v\n", solutions.LastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

	num := 19
	fmt.Printf("Number %d is a happy number: %t\n", num, solutions.IsHappy(num))

	digits := []int{9}
	fmt.Printf("Adding plus one equals: %v\n", solutions.PlusOne(digits))

	fmt.Printf("The value of myPow is %v\n", solutions.MyPow(2.00000, -2))
	fmt.Printf("The result of adding two string numbers is %s\n", solutions.AddStrings("289", "99"))
	fmt.Printf("The product of two string numbers is %s\n", solutions.Multiply("3866762897776739956", "15975363164662"))
	fmt.Printf("The single number in the array is %d\n", solutions.SingleNumber([]int{2, 2, 1}))
	fmt.Printf("The number is a power of two: %t\n", solutions.IsPowerOfTwo(0))

	numWeight, _ := strconv.ParseInt("00000000000000000000000010000000", 2, 64)
	fmt.Printf("The hamming weight of the number is %d\n", solutions.HammingWeight(uint32(numWeight)))

	fmt.Printf("The subsets of the array are %v\n", solutions.Subsets([]int{1, 2, 3}))
	fmt.Printf("The combination sum is %v\n", solutions.CombinationSum([]int{2, 3, 5}, 8))
	fmt.Printf("The permutation is %v\n", solutions.Permute([]int{1, 2, 3, 4}))
	fmt.Printf("The subsets with dupes are %v\n", solutions.SubsetsWithDup([]int{4, 4, 4, 1, 4}))

	fmt.Printf("The combination sum 2 is %v\n", solutions.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

	var board [][]byte = [][]byte{
		{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')},
	}
	word := "ABCCED"
	fmt.Printf("Word exists in grid: %t\n", solutions.Exist(board, word))

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
	fmt.Printf("The square root of x is %d\n", solutions.MySqrt(20))
	fmt.Printf("The number that i picked was %d\n", solutions.GuessNumber(10))
	fmt.Printf("The first bad version is %d\n", solutions.FirstBadVersion(10))
	fmt.Printf("The search insert position is %d\n", solutions.SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("The search range is %v\n", solutions.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Printf("The target is in search matrix %t\n", solutions.SearchMatrix(matrix, 3))
	fmt.Printf("The index of the element in the rotated array 2 is %t\n", solutions.RotatedSearch2([]int{1, 0, 1, 1, 1}, 0))
	fmt.Printf("The minimum element in the rotated array 2 is %d\n", solutions.FindMin2([]int{2, 2, 2, 0, 2}))
	fmt.Printf("The peak element in the array is %d\n", solutions.FindPeakElement([]int{1, 2, 3, 1}))
	fmt.Printf("The minimum size of subarray sum is %d\n", solutions.MinSubArrayLen(11, []int{1, 2, 3, 4, 5}))

	l1 := &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9}}}}}}}
	l2 := &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9, Next: &solutions.ListNode{Val: 9}}}}
	fmt.Printf("The sum of two numbers is %v\n", solutions.AddTwoNumbers(l1, l2))
	fmt.Printf("The number is a palindrome: %t\n", solutions.IsPalindromeNumber(12345654321))
	fmt.Printf("The integer after conversion is %d\n", solutions.MyAtoi("-2147483647"))
	fmt.Printf("The reverse of the integer is %d\n", solutions.Reverse(-123))
	fmt.Printf("Merging strings alternately yields %s\n", solutions.MergeAlternately("ab", "pqrs"))
	fmt.Printf("The greatest common divisor of strings is %s\n", solutions.GcdOfStrings("ABABAB", "ABAB"))
	fmt.Printf("The number of symmetric integers in the range is %d\n", solutions.CountSymmetricIntegers(1, 100))
	fmt.Printf("The minimum operations to make a special number are %d\n", solutions.MinimumOperations("19"))
	fmt.Printf("The strings can be made equal with operations: %t\n", solutions.CanBeEqual("bnxw", "bwxn"))
	fmt.Printf("The reverse vowels of a string a are %s\n", solutions.ReverseVowels("hello"))
	fmt.Printf("We can place flowers in the flowerbed: %t\n", solutions.CanPlaceFlowers([]int{1, 0, 1, 0, 0}, 1))
	fmt.Printf("The reverse words in a string are %s\n", solutions.ReverseWords("  hello world  "))

	arr := []int{1, 0, 0, 1, 0, 0, 2}
	solutions.MoveZeroes(arr)
	fmt.Printf("Moving Zeroes to the end yields %v\n", arr)

	fmt.Printf("Counting bits of n yields %v\n", solutions.CountBits(2))
	fmt.Printf("The extra characters in the string are %d\n", solutions.MinExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Printf("Converting roman to integer yields %d\n", solutions.RomanToInt("MCCCXCIX"))
	fmt.Printf("Converting integer to roman yields %s\n", solutions.IntToRoman(1399))

	rbl := &solutions.ListNode{Val: 1, Next: &solutions.ListNode{Val: 2, Next: &solutions.ListNode{Val: 3, Next: &solutions.ListNode{Val: 4, Next: &solutions.ListNode{Val: 5}}}}}
	val := solutions.ReverseBetween(rbl, 1, 4)
	fmt.Println(val.Next.Next.Next)
	fmt.Printf("Reversing the linked list yields %v\n", val)
	fmt.Printf("The result of the pascals triangle is %v\n", solutions.Generate(5))
	fmt.Printf("Is s a subsequence of t: %t\n", solutions.IsSubsequence("abc", "ahbgdc"))
	fmt.Printf("The combination sum 4 of the array of integers is %d\n", solutions.CombinationSum4([]int{1, 2, 3}, 4))
	fmt.Printf("The number of integer points covered with any part of the car are %d\n", solutions.NumberOfPoints([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Printf("Tha valid pickup and delivery options are %d\n", solutions.CountOrders(2))
	fmt.Printf("The result after grouping the people is %v\n", solutions.GroupThePeople([]int{2, 1, 3, 3, 3, 2}))
	fmt.Printf("The possible unique permutations that might contain duplicates are %v\n", solutions.PermuteUnique([]int{1, 1, 2}))
	fmt.Printf("The two strings can be made equal %t\n", solutions.CheckStrings("abcdba", "cabdab"))
	fmt.Printf("The maximum average subarray 1 is %v\n", solutions.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Printf("The minimum number of characters to delete are %d\n", solutions.MinDeletions("abcabc"))
	fmt.Printf("The maximum of almost unique subarrays is %d\n", solutions.MaxSum([]int{5, 9, 9, 2, 4, 5, 4}, 1, 3))
	fmt.Printf("The K closest elements of the array are %v\n", solutions.FindClosestElements([]int{1, 2, 3, 4, 5}, 3, 50))
	fmt.Printf("The number is a valid perfect square %t\n", solutions.IsPerfectSquare(14))
	fmt.Printf("The smallest letter greater than the target %s\n", string(solutions.NextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Printf("The minimum number of candies is %d\n", solutions.Candy([]int{1, 0, 2}))
	fmt.Printf("The intersection of two arrays is %d\n", solutions.Intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Printf("The intersection of two arrays ii is %d\n", solutions.Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Printf("The duplicate in the array is %d\n", solutions.FindDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("The Itinerary is %v\n", solutions.FindItinerary([][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	}))

	cycle := &solutions.ListNode{Val: 2, Next: &solutions.ListNode{Val: 0, Next: &solutions.ListNode{Val: 4}}}
	cycle.Next = cycle
	fmt.Printf("The linked list has a cycle: %t\n", solutions.HasCycle(&solutions.ListNode{Val: 3, Next: cycle}))
	fmt.Printf("The minimum cost to connect all points is %d\n", solutions.MinCostConnectPoints([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))
}

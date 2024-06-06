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
	g := methods.NewDirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)

	visitedOrder := make([]int, 0)
	g.DFS(g.Vertices[1], func(i int) {
		visitedOrder = append(visitedOrder, i)
	})

	fmt.Println("DFS", visitedOrder)

	visitedBFSOrder := make([]int, 0)
	g.BFS(g.Vertices[1], func(i int) {
		visitedBFSOrder = append(visitedOrder, i)
	})

	fmt.Println("BFS", visitedBFSOrder)

	fmt.Println("Printing all subsets of 3")
	utils.PrintAllSubsets(3, (1<<3)-1) // O(2^N * N)

	dp := [3][1 << 3]int{}
	for i := range dp {
		for k := range dp[i] {
			dp[i][k] = -1
		}
	}

	fmt.Printf("The maximum cost is %d\n", utils.MaximumCost(0, 0, 3, dp, [][]int{
		{2, 9, 7},
		{6, 4, 3},
		{1, 8, 5},
	})) // O(N * 2^N)

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

	head := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: nil}}
	fmt.Printf("The Reversed linked list is: %v\n", solutions.ReverseList(head))

	list1 := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 4, Next: nil}}}
	list2 := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: nil}}}
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

	reorderHead := &utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{
			Val: 3, Next: &utils.ListNode{Val: 4, Next: nil},
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

	l1 := &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9}}}}}}}
	l2 := &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9, Next: &utils.ListNode{Val: 9}}}}
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

	rbl := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}}}
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

	cycle := &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 0, Next: &utils.ListNode{Val: 4}}}
	cycle.Next = cycle
	fmt.Printf("The linked list has a cycle: %t\n", solutions.HasCycle(&utils.ListNode{Val: 3, Next: cycle}))
	fmt.Printf("The minimum cost to connect all points is %d\n", solutions.MinCostConnectPoints([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))
	fmt.Printf("There exists a triplet of indices i, j, k such that i < j < k: %t\n", solutions.IncreasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	fmt.Printf("The max number of k-sum pairs are %d\n", solutions.MaxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 4))
	fmt.Printf("The path with minimum effort is %d\n", solutions.MinimumEffortPath([][]int{
		{1, 2, 2}, {3, 8, 2}, {5, 3, 5},
	}))
	fmt.Printf("The shortest path visting all nodes is %d\n", solutions.ShortestPathLength([][]int{
		{1, 2, 3}, {0}, {0}, {0},
	}))
	fmt.Printf("The difference of two arrays is %v\n", solutions.FindDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
	fmt.Printf("The array has a unique number of occurences %t\n", solutions.UniqueOccurrences([]int{1, 2}))
	fmt.Printf("The minimum right shifts to sort the array are %d\n", solutions.MinimumRightShifts([]int{3, 4, 5, 1, 2}))
	fmt.Printf("The k weakest rows in a matrix are %v\n", solutions.KWeakestRows([][]int{
		{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1},
	}, 3))
	fmt.Printf("The sum of indices with k set bits are %d\n", solutions.SumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
	fmt.Printf("The possible way to select a group of students so that everyone remains happy is %d\n", solutions.CountWays([]int{6, 0, 3, 3, 6, 7, 2, 7}))
	fmt.Printf("The Binary tree Preorder Traversal of tree is %v\n", solutions.PreorderTraversal(&utils.TreeNode{
		Val: 1, Left: &utils.TreeNode{Val: 2}, Right: &utils.TreeNode{Val: 3},
	}))
	fmt.Printf("The Binary tree Inorder Traversal of tree is %v\n", solutions.InorderTraversal(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}},
	}))
	fmt.Printf("The Binary tree Postorder Traversal of tree is %v\n", solutions.PostorderTraversal(&utils.TreeNode{
		Val: 1, Left: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 6}}}, Right: &utils.TreeNode{Val: 7},
	}))
	fmt.Printf("The N array tree Postorder Traversal of tree is %v\n", solutions.Postorder(&utils.Node{
		Val: 1, Children: []*utils.Node{
			{Val: 2},
			{Val: 3, Children: []*utils.Node{{Val: 6}, {Val: 7, Children: []*utils.Node{{Val: 11}, {Val: 14}}}}},
			{Val: 4, Children: []*utils.Node{{Val: 8}, {Val: 12}}},
			{Val: 5, Children: []*utils.Node{{Val: 9, Children: []*utils.Node{{Val: 13}}}, {Val: 10}}},
		},
	}))
	fmt.Printf("The minimum number of operations to reduce x to zero are %d\n", solutions.MinOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Printf("The sub array sum is %d\n", solutions.SubarraySum([]int{-4, 3, 6, -2, 1, -1, 0, 2, -2, 3, 1}, 5))
	fmt.Printf("The pivot index of the array is %d\n", solutions.PivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Printf("The middle index in the array is %d\n", solutions.FindMiddleIndex([]int{2, 3, -1, 8, 4}))
	fmt.Printf("The head after removing the nth node from the end of the linked list is %v\n", solutions.RemoveNthFromEnd(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}},
	}, 2))
	fmt.Printf("The longest palindromic substring is %s\n", solutions.LongestPalindrome("babad"))
	fmt.Printf("The longest prefix is %s\n", solutions.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("The parentheses are %v\n", solutions.GenerateParenthesis(2))
	fmt.Printf("Swap nodes in pairs yields %v\n", solutions.SwapPairs(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4}}},
	}))
	fmt.Printf("The level order traversal is %v\n", solutions.LevelOrder(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}))
	fmt.Printf("The index of the string is %d\n", solutions.StrStr("sadbutsad", "sad"))
	fmt.Printf("Is valid BST %t\n", solutions.IsValidBST(&utils.TreeNode{
		Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3},
	}))
	fmt.Printf("Is palindrome linked list %t\n", solutions.IsPalindromeList(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 1}}},
	}))
	fmt.Printf("The longest string chain is %d\n", solutions.LongestStrChain([]string{"abcd", "dbqca"}))
	fmt.Printf("The longest increasing subsequence is %d\n", solutions.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Printf("The champagne tower Jth class in the ith row is this full %v\n", solutions.ChampagneTower(1, 1, 1))
	fmt.Printf("The maximum odd binary is %s\n", solutions.MaximumOddBinaryNumber("010"))
	fmt.Printf("The difference between strings s and t is %d\n", solutions.FindTheDifference("aeee", "aeeea"))
	fmt.Printf("The length of the last word is %d\n", solutions.LengthOfLastWord("Hello World"))
	fmt.Printf("Deleting duplicates from sorted list yields %v\n", solutions.DeleteDuplicates(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 1}}}},
	}))
	fmt.Printf("Deleting duplicates from sorted list 2 yields %v\n", solutions.DeleteDuplicates2(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{
			Val: 3, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{
				Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{
					Val: 4, Next: &utils.ListNode{Val: 5}},
				}},
			}}},
	}))
	fmt.Printf("Compressing the string yields %d\n", solutions.Compress([]byte{'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Printf("Removing duplicate letters yields %s\n", solutions.RemoveDuplicateLetters("bcabc"))
	fmt.Printf("The smallest subsequence of distinct characters is %s\n", solutions.SmallestSubsequence("cbacdcbc"))
	fmt.Printf("The maximum number of vowels in a substring is %d\n", solutions.MaxVowels("abciiidef", 3))
	fmt.Printf("The maximum consecutive ones are %d\n", solutions.FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Printf("The maximum consecutive ones wih k Os flipped is %d\n", solutions.LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Printf("The longest subarray of 1's after deleting one element is %d\n", solutions.LongestSubarray([]int{1, 1, 0, 1}))
	fmt.Printf("The decoded string at index is %s\n", solutions.DecodeAtIndex("y959q969u3hb22odq595", 222280369))
	fmt.Printf("The highest altitude is %d\n", solutions.LargestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Printf("The integer at rowIndex of the pascal's triangle is %v\n", solutions.GetRow(3))
	nums1 := []int{1, 2, 3, 0, 0, 0}
	solutions.Merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Printf("The merged sorted array is %d\n", nums1)
	fmt.Printf("Sorting an array by parity yields %v\n", solutions.SortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Printf("The longest substring in repeating character replaced array is %d\n", solutions.CharacterReplacement("ABAB", 2))
	fmt.Printf("The number of distinct solutions for the n-queens puzzle are, %d\n", solutions.TotalNQueens(4))
	fmt.Printf("The minimum number of operations to form a target array is %d\n", solutions.MinNumberOperations([]int{3, 1, 1, 2}))
	fmt.Printf("The sorted list is %v\n", solutions.MergeKLists([]*utils.ListNode{
		{Val: 1}, {Val: 2},
	}))
	fmt.Printf("Is the array monotonic %t\n", solutions.IsMonotonic([]int{-1, 3, 2}))
	fmt.Printf("The number of unique paths are %d\n", solutions.UniquePaths(3, 2))
	fmt.Println("The first unique character in the string is at position", solutions.FirstUniqChar("loveleetcode"))
	fmt.Printf("The maximum subarray sum is %d\n", solutions.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Printf("There is a 132 pattern in the array %t\n", solutions.Find132pattern([]int{1, 2, 3, 4}))
	fmt.Printf("The minimum cost climbing stairs is %d\n", solutions.MinCostClimbingStairs([]int{10, 15, 20}))

	RecentCounter := solutions.RecentConstructor()
	fmt.Println(RecentCounter.Ping(642))
	fmt.Println(RecentCounter.Ping(1849))
	fmt.Println(RecentCounter.Ping(4921))
	fmt.Println(RecentCounter.Ping(5936))
	fmt.Printf("The Number of recent calls are: %v\n", RecentCounter.Ping(5957))

	fmt.Printf("Deleting middle node of the linked list yields %v\n", solutions.DeleteMiddle(&utils.ListNode{
		Val: 2, Next: &utils.ListNode{Val: 1},
	}))
	fmt.Printf("The Odd even linked list is now %v\n", solutions.OddEvenList(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}},
	}))
	fmt.Printf("Reverse the words III yields %s\n", solutions.ReverseWordsIII("Let's take LeetCode contest"))
	fmt.Printf("The maximum triplet value of ordered triplet is %d\n", solutions.MaximumTripletValue([]int{12, 6, 1, 2, 7}))
	fmt.Printf("The maximum triplet value of ordered triplet II is %d\n", solutions.MaximumTripletValue2([]int{1, 10, 3, 4, 19}))
	fmt.Printf("The minimum size subarray in infinite array is %d\n", solutions.MinSizeSubarray([]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 5))
	fmt.Printf("The winner of the remove colored pieces if both neighbors are the same color is %t\n", solutions.WinnerOfGame("AAAABB"))
	fmt.Printf("The minimum operations to collect element are %d\n", solutions.MinCollectOperations([]int{3, 1, 5, 4, 2}, 2))
	fmt.Printf("The minimum number of operations to make array empty are %d\n", solutions.MinEmptyOperations([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}))
	nums2 := []int{2, 0, 2, 1, 1, 0}
	solutions.SortColors(nums2)
	fmt.Printf("Sorting the colors yields %v\n", nums2)
	fmt.Printf("The number of identical pairs in array are %d\n", solutions.NumIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Printf("The strings are close, %t\n", solutions.CloseStrings("aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff"))
	fmt.Printf("The result string after removing the stars is %s\n", solutions.RemoveStars("leet**cod*e"))
	fmt.Printf("The result of the array after asteroid collision is %v\n", solutions.AsteroidCollision([]int{5, 10, -5}))

	myHashMap := solutions.HashConstructor()
	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)
	fmt.Printf("The Value for key %d from the hashmap is %d\n", 1, myHashMap.Get(1))
	fmt.Printf("The Value for the key %d from the hashmap is %d\n", 3, myHashMap.Get(3))
	myHashMap.Put(2, 1)
	fmt.Printf("The Value for the key %d from the hashmap is %d\n", 2, myHashMap.Get(2))
	myHashMap.Remove(2)
	fmt.Printf("The Value for the key %d from the hashmap is %d\n", 2, myHashMap.Get(2))

	fmt.Printf("The tree leaves has a similar sequence %t\n", solutions.LeafSimilar(
		&utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 2}, Right: &utils.TreeNode{Val: 3}},
		&utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 2}},
	))
	fmt.Printf("The decoded string is %s\n", solutions.DecodeString("3[a2[c]]"))
	fmt.Printf("The majority element in the array is %d\n", solutions.MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Printf("The majority element 2 in the array is %d\n", solutions.MajorityElement2([]int{1, 1, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1}))
	fmt.Printf("The grid contains the same row column pairs %d\n", solutions.EqualPairs([][]int{
		{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2},
	}))
	fmt.Printf("The middle node of the linked list is %v\n", solutions.MiddleNode(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}},
	}))
	fmt.Printf("The twin pair sum of the linked list is %d\n", solutions.PairSum(&utils.ListNode{
		Val: 5, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 1}}},
	}))
	fmt.Printf("The maximum product after integer break is %d\n", solutions.IntegerBreak(10))
	fmt.Printf("Reversing nodes in k group yields %v\n", solutions.ReverseKGroup(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}},
	}, 2))
	fmt.Printf("The kth largest element in an array is %d\n", solutions.FindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Printf("The number of arrays that you can find the max k comparisons %v\n", solutions.NumOfArrays(2, 3, 1))

	inter := &utils.ListNode{Val: 8, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}
	fmt.Printf("The intersections of the two linked lists is %v\n", solutions.GetIntersectionNode(&utils.ListNode{
		Val: 4, Next: &utils.ListNode{Val: 1, Next: inter},
	}, &utils.ListNode{
		Val: 5, Next: &utils.ListNode{Val: 6, Next: &utils.ListNode{Val: 1, Next: inter}},
	}))
	fmt.Printf("The title to number of the excel colum is %d\n", solutions.TitleToNumber("AB"))
	fmt.Printf("The max dot product of two subsequences is %d\n", solutions.MaxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
	fmt.Printf("The difference of sums is %d\n", solutions.DifferenceOfSums(10, 3))
	fmt.Printf("The good nodes in the binary tree are %d\n", solutions.GoodNodes(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 3}}, Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 5}},
	}))
	fmt.Printf("The binary tree has path sum %t\n", solutions.HasPathSum(&utils.TreeNode{
		Val: 5, Left: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 11, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 2}}},
		Right: &utils.TreeNode{Val: 8, Left: &utils.TreeNode{Val: 13}, Right: &utils.TreeNode{Val: 4, Right: &utils.TreeNode{Val: 1}}},
	}, 22))
	fmt.Printf("The root-to-leaf paths in the binary tree are %v\n", solutions.PathSum2(&utils.TreeNode{
		Val: 5, Left: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 11, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 2}}},
		Right: &utils.TreeNode{Val: 8, Left: &utils.TreeNode{Val: 13}, Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 1}}},
	}, 22))
	fmt.Printf("The missing number in the array of distinct nums is %d\n", solutions.MissingNumber([]int{3, 0, 1}))
	fmt.Printf("Removing duplicates from the sorted array yields %d\n", solutions.RemoveDuplicates([]int{1, 1, 2}))
	fmt.Printf("The minimum operations to make array continous are %d\n", solutions.MinContinousOperations([]int{4, 8, 9, 11}))
	fmt.Printf("The paths in binary tree path sum 3 are %d\n", solutions.PathSum3(&utils.TreeNode{
		Val: 5, Left: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 11, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 2}}},
		Right: &utils.TreeNode{Val: 8, Left: &utils.TreeNode{Val: 13}, Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 1}}},
	}, 22))
	fmt.Printf("The node in the binary search tree is %v\n", solutions.SearchBST(&utils.TreeNode{
		Val:   4,
		Left:  &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3}},
		Right: &utils.TreeNode{Val: 7},
	}, 2))
	fmt.Printf("Removing elements from linked list yields %v\n", solutions.RemoveElements(&utils.ListNode{
		Val: 7, Next: &utils.ListNode{Val: 7, Next: &utils.ListNode{Val: 7, Next: &utils.ListNode{Val: 7}}},
	}, 7))
	fmt.Printf("The number of flowers that are full blown when the ith person arrives are %v\n", solutions.FullBloomFlowers([][]int{
		{1, 10}, {3, 3},
	}, []int{3, 3, 2}))
	fmt.Printf("The result afte rotating the array right k times is %v\n", solutions.RotateRight(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}},
	}, 2))
	fmt.Printf("The fibonacci of the number is %d\n", solutions.Fib(4))
	fmt.Printf("The peak index in the mountain array is %d\n", solutions.PeakIndexInMountainArray([]int{0, 2, 1, 0}))

	mountainArr := solutions.NewMountainArray([]int{1, 2, 3, 4, 5, 3, 1})
	fmt.Printf("Find in mountain array yields %d\n", solutions.FindInMountainArray(3, mountainArr))
	fmt.Printf("The minimum depth of the binary tree is %d\n", solutions.MinDepth(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}))
	fmt.Printf("Removing element from the array yields %d\n", solutions.RemoveElement([]int{3, 2, 2, 3}, 3))
	fmt.Printf("Deleting a node in the Binary search tree yields %v\n", solutions.DeleteNode(&utils.TreeNode{
		Val:   5,
		Left:  &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 2}, Right: &utils.TreeNode{Val: 4}},
		Right: &utils.TreeNode{Val: 6, Right: &utils.TreeNode{Val: 7}},
	}, 3))
	fmt.Printf("The longest consecutive sequence is %d\n", solutions.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Printf("The maximum level sum of a binary tree is %d\n", solutions.MaxLevelSum(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 7, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: -8}},
		Right: &utils.TreeNode{Val: 0},
	}))
	fmt.Printf("The binary tree right side view is %v\n", solutions.RightSideView(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 5}},
		Right: &utils.TreeNode{Val: 3, Right: &utils.TreeNode{Val: 4}},
	}))
	fmt.Printf("The minimum amount of money required to paint n walls is %d\n", solutions.PaintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
	mat := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	solutions.SetZeroes(mat)
	fmt.Printf("The matrix after setting zeros in place is %v\n", mat)

	bit32, _ := strconv.ParseInt("00000010100101000001111010011100", 2, 32)
	fmt.Printf("Reversing bits yields %d\n", solutions.ReverseBits(uint32(bit32)))
	fmt.Printf("The number of ways to reach a position after exactly k steps is %d\n", solutions.NumberOfWays(1, 2, 3))
	fmt.Printf("The number of ways to stay in the same place after some steps is %d\n", solutions.NumWays(3, 2))
	fmt.Printf("The number of complete tree nodes are %d\n", solutions.CountNodes(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 4}, Right: &utils.TreeNode{Val: 5}},
		Right: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 6}},
	}))
	fmt.Printf("The indices wih index and value difference 1 are %v\n", solutions.FindIndices([]int{5, 1, 4, 1}, 2, 4))
	fmt.Printf("The tribonacci of n is %d\n", solutions.Tribonacci(4))
	fmt.Printf("The minimum path sum from top to bottom is %d\n", solutions.MinimumTotal([][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}))
	fmt.Printf("The minimum path sum of the m*n grid is %d\n", solutions.MinPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))
	fmt.Printf("The nodes form a valid binary tree %t\n", solutions.ValidateBinaryTreeNodes(4,
		[]int{1, -1, 3, -1}, []int{2, -1, -1, -1},
	))
	fmt.Printf("The fewest number of coins needed to make up the amount are %d\n", solutions.CoinChange([]int{1, 2, 5}, 11))
	fmt.Printf("The array contains nearby duplicate values %t\n", solutions.ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Printf("The minimum number of months needed to complete all the courses is %d\n", solutions.MinimumTime(3, [][]int{
		{1, 3}, {2, 3},
	}, []int{3, 2, 5}))
	fmt.Printf("The number of possible unique paths that a robot can make are %d\n", solutions.UniquePathsWithObstacles([][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}))
	fmt.Printf("The integer returned as a string is %v\n", solutions.FizzBuzz(5))
	fmt.Printf("The lowercase version of the string is %s\n", solutions.ToLowerCase("LOVELY"))
	fmt.Printf("The strings are equal when both typed into text editors %t\n", solutions.BackspaceCompare("ab#c", "ad#c"))

	mat2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	solutions.Rotate(mat2)
	fmt.Printf("Rotating the matrix yields %v\n", mat2)
	fmt.Printf("The spiral order of the matrix is %v\n", solutions.SpiralOrder([][]int{
		{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30},
	}))
	fmt.Printf("The transpose of the matrix is %v\n", solutions.Transpose([][]int{
		{1, 2, 3}, {4, 5, 6},
	}))
	fmt.Printf("Reversing only letters yields %s\n", solutions.ReverseOnlyLetters("Test1ng-Leet=code-Q!"))
	fmt.Printf("After Merging intervals the array is %v\n", solutions.MergeIntervals([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	fmt.Printf("The array is sorted and rotated %t\n", solutions.Check([]int{3, 4, 5, 1, 2}))
	fmt.Printf("The intervals after insertion are %v\n", solutions.Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Printf("Generating the matrix in spiral order yields %v\n", solutions.GenerateMatrix(3))
	fmt.Printf("Removing duplicates from the sorted array II yields %d\n", solutions.RemoveDuplicates2([]int{1, 1, 1, 2, 2, 3}))

	cycle2 := &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 0, Next: &utils.ListNode{Val: -4}}}
	cycle2.Next = cycle2
	fmt.Printf("There is a cycle in the linked list: %v\n", solutions.DetectCycle(&utils.ListNode{Val: 3, Next: cycle2}))
	fmt.Printf("The non overlapping intervals are %v\n", solutions.EraseOverlapIntervals([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}))
	fmt.Printf("The constrained subsequence sum is %d\n", solutions.ConstrainedSubsetSum([]int{
		10, 2, -10, 5, 20,
	}, 2))
	fmt.Printf("The sliding window maximum is %v\n", solutions.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("The strings are isomorphic %t\n", solutions.IsIsomorphic("egg", "add"))
	fmt.Printf("The result after repeatedly adding all its digits is %d\n", solutions.AddDigits(38))
	fmt.Printf("The third maximum number in the array is %d\n", solutions.ThirdMax([]int{3, 2, 1}))
	fmt.Printf("The maximum score of a good subarray in the array is %d\n", solutions.MaximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
	fmt.Printf("The minimum sum of mountain triplets I is %d\n", solutions.MinimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Printf("The minimum sum of mountain triplets II is %d\n", solutions.MinimumSum2([]int{5, 4, 8, 7, 10, 2}))
	fmt.Printf("String s follows pattern %t\n", solutions.WordPattern("abba", "dog cat cat dog"))
	fmt.Printf("Number is a power of 4 %t\n", solutions.IsPowerOfFour(16))
	fmt.Printf("Partitioning the array based on a pivot yields %v\n", solutions.PivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
	fmt.Printf("Re arranging array elements by sign yields %v\n", solutions.RearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	fmt.Printf("Sorting array by parity II yields %v\n", solutions.SortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Printf("Partioning the list yields %v\n", solutions.PartitionList(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 5, Next: &utils.ListNode{Val: 2}}}}},
	}, 3))
	fmt.Printf("The largest value for each tree row is %v\n", solutions.LargestValues(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 3}},
		Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 9}},
	}))
	fmt.Printf("Sorting an array in ascending order yields %v\n", solutions.SortArray([]int{5, 2, 3, 1}))
	fmt.Printf("The concatenation of two nums arrays is %v\n", solutions.GetConcatenation([]int{1, 2, 1}))
	fmt.Printf("The string is a valid palindrom after deleting atmost one character from it %t\n", solutions.ValidPalindrome("abca"))
	fmt.Printf("Shifting 2d grid yields %v\n", solutions.ShiftGrid([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}, 1))
	fmt.Printf("The kth symbol in grammar is %d\n", solutions.KthGrammar(3, 2))
	fmt.Printf("Replacing elements with greatest element on the right side equals %v\n", solutions.ReplaceElements([]int{17, 18, 5, 4, 6, 1}))

	rotateArr := []int{1, 2, 3, 4, 5, 6, 7}
	solutions.RotateArray(rotateArr, 3)
	fmt.Printf("Rotating the array k times yields %v\n", rotateArr)
	fmt.Printf("All numbers disappeared from array are %v\n", solutions.FindDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Printf("Number is a power of 3 %t\n", solutions.IsPowerOfThree(27))
	fmt.Printf("The number of factored binary trees are %d\n", solutions.NumFactoredBinaryTrees([]int{2, 4, 5, 10}))
	fmt.Printf("The result after soting the list is %v\n", solutions.SortList(&utils.ListNode{
		Val: 4, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 3}}},
	}))
	fmt.Printf("The result after swapping nodes in the list is %v\n", solutions.SwapNodes(utils.NewListNode([]int{1, 2, 3, 4, 5}), 2))
	fmt.Printf("The sign of the product array is %d\n", solutions.ArraySign([]int{-1, -2, -3, -4, 3, 2, 1}))

	parkingSystem := solutions.ParkingConstructor(1, 1, 0)
	fmt.Println("There is 1 available slot for a big car", parkingSystem.AddCar(1))
	fmt.Println("There is 1 available slot for a medium car", parkingSystem.AddCar(2))
	fmt.Println("There is no available slot for a small car", parkingSystem.AddCar(3))
	fmt.Println("There is no available slot for a big car. It is already occupied", parkingSystem.AddCar(1))

	fmt.Printf("The number of palindromic substrings are %d\n", solutions.CountSubstrings("aaa"))

	minStack := solutions.MinStackConstructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("The current min value is %d\n", minStack.GetMin())
	minStack.Pop()
	fmt.Printf("The current top value is %d\n", minStack.Top())
	fmt.Printf("The current min value is %d\n", minStack.GetMin())

	fmt.Printf("The result after evaluating the expression is %d\n", solutions.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Printf("The number of ways you can make up the amount from the coins is %d\n", solutions.Change(5, []int{1, 2, 5}))
	fmt.Printf("The number of strings that can be formed are %d\n", solutions.CountVowelPermutation(2))
	fmt.Printf("The sum of scores on the record after applying all the operations is %d\n", solutions.CalPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Printf("The number of unique email addresses are %d\n", solutions.NumUniqueEmails([]string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}))

	numArray := solutions.SumRangeConstructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("The sum of elements of nums between indices 0 and 2 are %d\n", numArray.SumRange(0, 2))
	fmt.Printf("The sum of elements of nums between indices 2 and 5 are %d\n", numArray.SumRange(2, 5))
	fmt.Printf("The sum of elements of nums between indices 0 and 5 are %d\n", numArray.SumRange(0, 5))

	fmt.Printf("The minimum number of pigs needed to figure out which bucket is poisonous are %d\n", solutions.PoorPigs(4, 15, 15))
	fmt.Printf("The possible combinations of k numbers chosen from 1-n are %v\n", solutions.Combine(4, 2))
	fmt.Printf("The kth smallest element in the BST is %d\n", solutions.KthSmallest(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}}, Right: &utils.TreeNode{Val: 4},
	}, 1))
	fmt.Printf("The head of the modified linked list after removing elements is %v\n", solutions.RemoveNodes(&utils.ListNode{
		Val: 5, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 13, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 8}}}},
	}))

	node := &utils.ListNode{Val: 5, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 9}}}
	deleteNode := &utils.ListNode{Val: 4, Next: node}
	solutions.DeleteNode2(node)
	fmt.Printf("After deleting the node, the list is %v\n", deleteNode)

	fmt.Printf("Sorting integers by number of 1 bits yields %v\n", solutions.SortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	fmt.Printf("The minimum difference between highest and lowest of k scores is %d\n", solutions.MinimumDifference([]int{9, 4, 1, 7}, 2))
	fmt.Printf("The number of odd numbers in an interval range are %d\n", solutions.CountOdds(8, 10))
	fmt.Printf("The matrix diagonal sum is %d\n", solutions.DiagonalSum([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
	fmt.Printf("The squares of a sorted array are %v\n", solutions.SortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Printf("The array that satisfies xor is %v\n", solutions.FindArray([]int{5, 2, 0, 3, 1}))
	fmt.Printf("The maximum number of instances of balloon that can be formed are %d\n", solutions.MaxNumberOfBalloons("loonbalxballpoon"))
	fmt.Printf("The sentence is a pangram %t\n", solutions.CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Printf("Sorting the people in descending order by people's heights yields %v\n", solutions.SortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Printf("The modes in the binary search tree are %v\n", solutions.FindMode(&utils.TreeNode{
		Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 2},
	}))
	fmt.Printf("Transforming the arrasy by rank yields %v\n", solutions.ArrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Printf("The party that will finally anounce victory is %s\n", solutions.PredictPartyVictory("RDD"))

	trie := solutions.TrieConstructor()
	trie.Insert("apple")
	fmt.Println("Search apple", trie.Search("apple"))
	fmt.Println("Search app", trie.Search("app"))
	fmt.Println("Starts with app", trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println("Search app", trie.Search("app"))

	fmt.Printf("The number of nodes equal to equal to average of the subtree are %d\n", solutions.AverageOfSubtree(&utils.TreeNode{
		Val:   4,
		Left:  &utils.TreeNode{Val: 8, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 1}},
		Right: &utils.TreeNode{Val: 5, Right: &utils.TreeNode{Val: 6}},
	}))
	fmt.Printf("Shuffling the array yields %v\n", solutions.Shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Printf("Adding array-form of integer yields %v\n", solutions.AddToArrayForm([]int{2, 1, 5}, 806))
	fmt.Printf("The sum of multiples in range are %d\n", solutions.SumOfMultiples(7))
	fmt.Printf("The stack operations needed to build target are %v\n", solutions.BuildArray([]int{1, 3}, 3))
	fmt.Printf("Removing all adjacent duplicates in string yields %s\n", solutions.RemoveDuplicates3("abbaca"))
	fmt.Printf("Removing all adjacent k duplicates in string yields %s\n", solutions.RemoveDuplicates4("deeedbbcccbdaa", 3))
	fmt.Printf("The array ans is %v\n", solutions.BuildArray2([]int{0, 2, 1, 5, 3, 4}))
	fmt.Printf("Last moment before all ants fall out of a plank is %d\n", solutions.GetLastMoment(4, []int{4, 3}, []int{0, 1}))
	fmt.Printf("Sorting the sentence yields %v\n", solutions.SortSentence("is2 sentence4 This1 a3"))
	fmt.Printf("The number of good pairs with absolute difference k are %d\n", solutions.CountKDifference([]int{1, 2, 2, 1}, 1))
	fmt.Printf("Flipping and inverting an image yields %v\n", solutions.FlipAndInvertImage([][]int{
		{1, 1, 0}, {1, 0, 1}, {0, 0, 0},
	}))
	fmt.Printf("The winner of an array game is %d\n", solutions.GetWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))

	p := &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 6}, Right: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 4}}}
	q := &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 8}}
	fmt.Printf("The lowest common ancestor of a binary tree is %v\n", solutions.LowestCommonAncestor(&utils.TreeNode{Val: 3, Left: p, Right: q}, p, q))
	fmt.Printf("The hamming distance of integers x and y is %d\n", solutions.HammingDistance(1, 4))
	fmt.Printf("The minimum bit flips to convert number %d\n", solutions.MinBitFlips(10, 7))
	fmt.Printf("The minimum flips to make a OR b equal to c is %d\n", solutions.MinFlips(2, 6, 5))

	seatManager := solutions.ReserveConstructor(5)
	fmt.Println("Reserving", seatManager.Reserve())
	fmt.Println("Reserving", seatManager.Reserve())
	seatManager.Unreserve(2)
	fmt.Println("Reserving", seatManager.Reserve())
	fmt.Println("Reserving", seatManager.Reserve())
	fmt.Println("Reserving", seatManager.Reserve())
	fmt.Println("Reserving", seatManager.Reserve())
	seatManager.Unreserve(5)

	fmt.Printf("The minimum number of arrows to burst balloons is %d\n", solutions.FindMinArrowShots([][]int{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}))
	fmt.Printf("The number of days to wait after ith day to get warmer temperature are %v\n", solutions.DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Printf("The next greater element I is %v\n", solutions.NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Printf("The maximum number of monsters to eliminate are %d\n", solutions.EliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	fmt.Printf("Inserting greatest common divisor in linked list yields %v\n", solutions.InsertGreatestCommonDivisors(&utils.ListNode{
		Val: 18, Next: &utils.ListNode{Val: 6, Next: &utils.ListNode{Val: 10, Next: &utils.ListNode{Val: 3}}},
	}))
	fmt.Printf("The deepest leaves sum is %d\n", solutions.DeepestLeavesSum(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 7}}, Right: &utils.TreeNode{Val: 5}},
		Right: &utils.TreeNode{Val: 3, Right: &utils.TreeNode{Val: 6, Right: &utils.TreeNode{Val: 8}}},
	}))
	fmt.Printf("The 2D array from array is %v\n", solutions.FindMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Printf("The cell is reachable after t seconds %t\n", solutions.IsReachableAtTime(2, 4, 7, 7, 6))
	fmt.Printf("Converting BST to greater tree yields %v\n", solutions.ConvertBST(&utils.TreeNode{
		Val:   4,
		Left:  &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}}},
		Right: &utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7, Right: &utils.TreeNode{Val: 8}}},
	}))
	fmt.Printf("Converting BST to greater tree yields %v\n", solutions.BstToGst(&utils.TreeNode{
		Val:   4,
		Left:  &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}}},
		Right: &utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}, Right: &utils.TreeNode{Val: 7, Right: &utils.TreeNode{Val: 8}}},
	}))
	fmt.Printf("The xor operation of the array is %d\n", solutions.XorOperation(5, 0))
	fmt.Printf("he max power of the string is %d\n", solutions.MaxPower("abbcccddddeeeeedcba"))
	fmt.Printf("The longest continous increasing subsequence is %d\n", solutions.FindLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Printf("The number of homogenous substrings is %d\n", solutions.CountHomogenous("zzzzz"))
	fmt.Printf("The number of substrings with onl 1s are %d\n", solutions.NumSub("111111"))
	fmt.Printf("The number of zero filled subarrays is %d\n", solutions.ZeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	fmt.Printf("The longest contiguous segment of 1s is strictly longer than the zeroes %t\n", solutions.CheckZeroOnes("1101"))
	fmt.Printf("Restoring the array from adjacent pairs yields %v\n", solutions.RestoreArray([][]int{
		{2, 1}, {3, 4}, {3, 2},
	}))
	fmt.Printf("The integer n is strictly palindromic %t\n", solutions.IsStrictlyPalindromic(9))
	fmt.Printf("Merging nodes in between zeros yields %v\n", solutions.MergeNodes(utils.NewListNode([]int{0, 3, 1, 0, 4, 5, 2, 0})))
	fmt.Printf("The sum of nodes with even valued grandparents is %d\n", solutions.SumEvenGrandparent(&utils.TreeNode{
		Val:   6,
		Left:  &utils.TreeNode{Val: 7, Left: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 9}}, Right: &utils.TreeNode{Val: 7, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 4}}},
		Right: &utils.TreeNode{Val: 8, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3, Right: &utils.TreeNode{Val: 5}}},
	}))

	graph := solutions.GraphConstructor(4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})
	fmt.Printf("The cost of the shortest path from 3 to 2 is %d\n", graph.ShortestPath(3, 2))
	fmt.Printf("The cost of the shortest path from 0 to 3 is %d\n", graph.ShortestPath(0, 3))
	graph.AddEdge([]int{1, 3, 4})
	fmt.Printf("The cost of the shortest path from 0 to 3 is %d\n", graph.ShortestPath(0, 3))

	fmt.Printf("The least number of buses you must take to travel from sourc to target are %d\n", solutions.NumBusesToDestination([][]int{
		{1, 2, 7}, {3, 6, 7},
	}, 1, 6))
	fmt.Printf("Sorting the vowels in a string yields %s\n", solutions.SortVowels("LQRamBOHfq"))
	fmt.Printf("The triangular sum of an array is %d\n", solutions.TriangularSum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("The minimum number of steps to make two strings anagram is %d\n", solutions.MinSteps("leetcode", "practice"))
	fmt.Printf("The minimum number of steps to make two strings anagram II is %d\n", solutions.MinSteps2("leetcode", "coats"))
	fmt.Printf("The Number of distinct integers after reverse operations is %d\n", solutions.CountDistinctIntegers([]int{1, 13, 10, 12, 31}))
	fmt.Printf("The number of unique palindromic subsequences of length three are %d\n", solutions.CountPalindromicSubsequence("aabca"))
	fmt.Printf("The minimum number of steps to make word1 and word2 the same is %d\n", solutions.MinDistance("sea", "eat"))
	fmt.Printf("The result of s after removing all occurrences of a substring is %s\n", solutions.RemoveOccurrences("daabcbaabcbc", "abc"))
	fmt.Printf("The duplicates in the arraya are %v\n", solutions.FindDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Printf("The maximum possible value of an element after decrementing and rearranging is %d\n", solutions.MaximumElementAfterDecrementingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Printf("The level order traversal of an n-array tree is %v\n", solutions.NArryLevelOrder(&utils.Node{
		Val: 1, Children: []*utils.Node{
			{Val: 3, Children: []*utils.Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4},
		},
	}))
	fmt.Printf("The level order traversal II is %v\n", solutions.LevelOrderBottom(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}))
	fmt.Printf("The Zigzag level order traversal is %v\n", solutions.ZigzagLevelOrder(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}))
	fmt.Printf("Adding two integers yields %d\n", solutions.Sum(12, 5))
	fmt.Printf("The unique binary string is %s\n", solutions.FindDifferentBinaryString([]string{"01", "10"}))
	fmt.Printf("The single non duplicate element in the sorted array is %d\n", solutions.SingleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Printf("Coverting binary number in a linked list to integer yields %d\n", solutions.GetDecimalValue(&utils.ListNode{
		Val: 1, Next: &utils.ListNode{Val: 0, Next: &utils.ListNode{Val: 1}},
	}))
	fmt.Printf("The minimized max pair sum in an array is %d\n", solutions.MinPairSum([]int{3, 5, 2, 3}))
	fmt.Printf("The single number II in the array is %d\n", solutions.SingleNumber2([]int{2, 2, 3, 2}))
	fmt.Printf("The single numbers II in the array are %v\n", solutions.SingleNumber3([]int{1, 2, 1, 3, 2, 5}))
	fmt.Printf("The sum of left leaves is %d\n", solutions.SumOfLeftLeaves(&utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 9},
		Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}))
	fmt.Printf("The running sum of a 1d array is %v\n", solutions.RunningSum([]int{1, 2, 3, 4}))
	fmt.Printf("The lonely numbers in the array are %v\n", solutions.FindLonely([]int{10, 6, 5, 8}))
	fmt.Printf("The maximum possible frequency of an element after performing k operations is %d\n", solutions.MaxFrequency([]int{1, 2, 4}, 5))
	fmt.Printf("The minimum number of positive deci-binary numbers needed are %d\n", solutions.MinPartitions("82734"))
	fmt.Printf("The tree constructed from preorder traversal is %v\n", solutions.BstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
	fmt.Printf("The number of pairs whose sum is less than target is %d\n", solutions.CountPairs([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Printf("The number of operations to make all elements in nums equal is %d\n", solutions.ReductionOperations([]int{5, 1, 3}))
	fmt.Printf("Merging linked list in between yields %v\n", solutions.MergeInBetween(&utils.ListNode{
		Val: 0, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: &utils.ListNode{Val: 5}}}}},
	}, 3, 4, &utils.ListNode{Val: 1000000, Next: &utils.ListNode{Val: 1000001, Next: &utils.ListNode{Val: 1000002}}}))
	fmt.Printf("Reversing odd levels yields %v\n", solutions.ReverseOddLevels(&utils.TreeNode{
		Val:   2,
		Left:  &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 8}, Right: &utils.TreeNode{Val: 13}},
		Right: &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 21}, Right: &utils.TreeNode{Val: 34}},
	}))
	fmt.Printf("The left annd right sum difference is %v\n", solutions.LeftRightDifference([]int{10, 4, 8, 3}))
	fmt.Printf("The two arrays represent the same string %t\n", solutions.ArrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Printf("The minimum amount of time to collect the garbage is %d\n", solutions.GarbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
	fmt.Printf("The maximum product subarray is %d\n", solutions.MaxProduct([]int{2, 3, -2, 4}))
	fmt.Printf("The maximum product of three numbers is %d\n", solutions.MaximumProduct([]int{-1, -2, -3}))
	fmt.Printf("The total sum of root to leaf numbers is %d\n", solutions.SumNumbers(&utils.TreeNode{
		Val: 1, Left: &utils.TreeNode{Val: 2}, Right: &utils.TreeNode{Val: 3},
	}))
	fmt.Printf("The prefix common array of two arrays is %v\n", solutions.FindThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	fmt.Printf("The number of nice pairs in the array is %d\n", solutions.CountNicePairs([]int{42, 11, 1, 97}))
	fmt.Printf("The number of bad pairs are %d\n", solutions.CountBadPairs([]int{4, 1, 3, 3}))
	fmt.Printf("The k-diff pairs in the array are %d\n", solutions.FindPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Printf("The number of equal and divisible pairs of an array are %d\n", solutions.CountPairs2([]int{3, 1, 2, 2, 2, 1, 3}, 2))
	fmt.Printf("The maximum number of words found in sentences are %d\n", solutions.MostWordsFound([]string{
		"alice and bob love leetcode", "i think so too", "this is great thanks very much",
	}))
	fmt.Printf("The diagonal order II of the elements in the 2d array is %v\n", solutions.FindDiagonalOrder2([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
	fmt.Printf("The result after sorting students by their kth score is %v\n", solutions.SortTheStudents([][]int{
		{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15},
	}, 2))
	fmt.Printf("The permutation that satisfies the property is %s\n", solutions.CustomSortString("cba", "abcd"))
	fmt.Printf("Sorting characters by frequency yields %s\n", solutions.FrequencySort("Aabb"))
	fmt.Printf("Sorting the array by frequency yields %v\n", solutions.FrequencySort2([]int{1, 1, 2, 2, 2, 3}))
	fmt.Printf("The array can be re arranged to form an arithmetic progression %t\n", solutions.CanMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Printf("The number of arithmetic subarrays are %d\n", solutions.NumberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Printf("The arithmetic subarrays are %v\n", solutions.CheckArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Printf("The length of the longest alphabetical continous substring is %d\n", solutions.LongestContinuousSubstring("abacaba"))
	fmt.Printf("The target array is %v\n", solutions.CreateTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Printf("The maximum number of coins you can get are %d\n", solutions.MaxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Printf("The Bottm left tree value %d\n", solutions.FindBottomLeftValue(&utils.TreeNode{
		Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3},
	}))
	fmt.Printf("Inserting into BST yields %v\n", solutions.InsertIntoBST(&utils.TreeNode{
		Val: 4, Left: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 3}}, Right: &utils.TreeNode{Val: 7},
	}, 5))
	fmt.Printf("Truncating the senetence yields %s\n", solutions.TruncateSentence("Hello how are you Contestant", 4))
	fmt.Printf("The root equal sum of children %t\n", solutions.CheckTree(&utils.TreeNode{
		Val: 10, Left: &utils.TreeNode{Val: 4}, Right: &utils.TreeNode{Val: 6},
	}))
	fmt.Printf("The sum of the absolute differences in a sorted array is %v\n", solutions.GetSumAbsoluteDifferences([]int{2, 3, 5}))
	fmt.Printf("The tree is even odd %t\n", solutions.IsEvenOddTree(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 10, Left: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 12}, Right: &utils.TreeNode{Val: 8}}},
		Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 7, Left: &utils.TreeNode{Val: 6}}, Right: &utils.TreeNode{Val: 9, Right: &utils.TreeNode{Val: 2}}},
	}))
	fmt.Printf("The diagonal order of the elements in a matrix are %v\n", solutions.FindDiagonalOrder([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
	fmt.Printf("The result after restoring string is %s\n", solutions.RestoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Printf("The difference between the element sum and the digit sum is %d\n", solutions.DifferenceOfSum([]int{1, 15, 6, 3}))
	fmt.Printf("The largest submatrix with rearrangements is %v\n", solutions.LargestSubmatrix([][]int{
		{0, 0, 1}, {1, 1, 1}, {1, 0, 1},
	}))
	fmt.Printf("The perimeter of the island is %d\n", solutions.IslandPerimeter([][]int{
		{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
	}))
	fmt.Printf("The Maximum binary tree is %v\n", solutions.ConstructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	fmt.Printf("The words that match the pattern are %v\n", solutions.FindAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Printf("The maximum wealth is %d\n", solutions.MaximumWealth([][]int{
		{1, 2, 3}, {3, 2, 1},
	}))
	fmt.Printf("The number of distinct phone numbers of length n we can dial are %d\n", solutions.KnightDialer(2))
	fmt.Printf("The diagonal order of the matrix is %v\n", solutions.DiagonalSort([][]int{
		{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2},
	}))
	fmt.Printf("The number of possible non empty tile possibilites are %d\n", solutions.NumTilePossibilities("AAB"))
	fmt.Printf("The new text after arranging the words is %s\n", solutions.ArrangeWords("Leetcode is cool"))
	fmt.Printf("The indices of the words containing character x are %v\n", solutions.FindWordsContaining([]string{"leet", "code"}, 'e'))
	fmt.Printf("The number of ways to split array are %d\n", solutions.WaysToSplitArray([]int{10, 4, -8, 7}))
	fmt.Printf("The number of ways to divide a long corridor are %d\n", solutions.NumberOfWays2("SSPPSPS"))
	fmt.Printf("The max sum of an hour glass is %d\n", solutions.MaxSum2([][]int{
		{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9},
	}))
	fmt.Printf("The xor queries of a subarray are %v\n", solutions.XorQueries([]int{1, 3, 4, 8}, [][]int{
		{0, 1}, {1, 2}, {0, 3}, {3, 3},
	}))
	fmt.Printf("The sum of odd length subarrays is %d\n", solutions.SumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Printf("The total hamming distance between all pairs of integers is %d\n", solutions.TotalHammingDistance([]int{4, 14, 2}))
	fmt.Printf("Adding binary strings yields %s\n", solutions.AddBinary("11", "1"))
	fmt.Printf("The maximum product of word lengths is %d\n", solutions.MaxProduct2([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Printf("The number of islands is %d\n", solutions.NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	fmt.Printf("The number of digits that divide the number are %d\n", solutions.CountDigits(1248))
	fmt.Printf("The minimum one bit operations to make integer zero is %d\n", solutions.MinimumOneBitOperations(3))
	fmt.Printf("You can visit all rooms %t\n", solutions.CanVisitAllRooms([][]int{
		{1}, {2}, {3}, {},
	}))
	fmt.Printf("The number of provinces are %d\n", solutions.FindCircleNum([][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}))
	fmt.Printf("The height balanced binary search tree from array is %v\n", solutions.SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Printf("The height balanced binary search tree from list is %v\n", solutions.SortedListToBST(&utils.ListNode{
		Val: -10, Next: &utils.ListNode{Val: -3, Next: &utils.ListNode{Val: 0, Next: &utils.ListNode{Val: 5, Next: &utils.ListNode{Val: 9}}}},
	}))
	fmt.Printf("The nth term of the count and say sequence is %s\n", solutions.CountAndSay(4))
	fmt.Printf("The resulting string is %s\n", solutions.Evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}))
	fmt.Printf("The score of the parentheses is %d\n", solutions.ScoreOfParentheses("(()(()))"))
	fmt.Printf("The resulting string after reversing the prefix is %s\n", solutions.ReversePrefix("abcdefd", 'd'))
	fmt.Printf("The first palindromic string in the array %s\n", solutions.FirstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
	fmt.Printf("The words that can be formed by characters are %d\n", solutions.CountCharacters([]string{"cat", "bt", "hat", "tree"}, "attach"))
	fmt.Printf("The maximum number of copies of target that can be formed are %d\n", solutions.RearrangeCharacters("ilovecodingonleetcode", "code"))
	fmt.Printf("The quotient after dividing two integers %d\n", solutions.Divide(10, 3))
	fmt.Printf("The array could become non decreasing %t\n", solutions.CheckPossibility([]int{4, 2, 3}))
	fmt.Printf("The defanged version of the IP address is %s\n", solutions.DefangIPaddr("1.1.1.1"))
	fmt.Printf("The minimum time visiting all points is %d\n", solutions.MinTimeToVisitAllPoints([][]int{
		{1, 1}, {3, 4}, {-1, 0},
	}))
	fmt.Printf("The smallest value of the rearranged number is %d\n", solutions.SmallestNumber(310))
	fmt.Printf("The usage of the capitals is right %t\n", solutions.DetectCapitalUse("USA"))
	fmt.Printf("The largest number you can form is %s\n", solutions.LargestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Printf("The capitalized tite is %s\n", solutions.CapitalizeTitle("capiTalIze tHe titLe"))
	fmt.Printf("The maximum good integer in string is %s\n", solutions.LargestGoodInteger("6777133339"))

	subrectangleQueries := solutions.SubrectangleQueriesConstructor([][]int{
		{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1},
	})
	fmt.Printf("Get value at row 0, col 2 %d\n", subrectangleQueries.GetValue(0, 2))
	subrectangleQueries.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Printf("Get value at row 0, col 2 %d\n", subrectangleQueries.GetValue(0, 2))
	fmt.Printf("Get value at row 3, col 1 %d\n", subrectangleQueries.GetValue(3, 1))
	fmt.Printf("The maximum total sum to keep city skyline %d\n", solutions.MaxIncreaseKeepingSkyline([][]int{
		{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0},
	}))
	fmt.Printf("The 2D array representing the matrix requirements %v\n", solutions.RestoreMatrix([]int{3, 8}, []int{4, 7}))
	fmt.Printf("The largest odd number is %s\n", solutions.LargestOddNumber("52"))
	fmt.Printf("The number of matches played in the tournament until winner is decided are %d\n", solutions.NumberOfMatches(7))
	fmt.Printf("S is a valid string %t\n", solutions.IsValid2("aabcbc"))
	fmt.Printf("Converting non negative integer num to english words yields %s\n", solutions.NumberToWords(123))
	fmt.Printf("Printing strings vertically yields %v\n", solutions.PrintVertically("TO BE OR NOT TO BE"))
	fmt.Printf("The number of stones that are also jewels are %d\n", solutions.NumJewelsInStones("aA", "aAAbbbb"))
	fmt.Printf("The total amount of money he will have in leetcode bank is %d\n", solutions.TotalMoney(20))
	fmt.Printf("The number of trailing zeros in n! is %d\n", solutions.TrailingZeroes(5))
	fmt.Printf("The number of primes is %d\n", solutions.CountPrimes(10))
	fmt.Printf("The number is an ugly number %t\n", solutions.IsUgly(6))
	fmt.Printf("Add convert the temperature %v\n", solutions.ConvertTemperature(36.50))
	fmt.Printf("The target exist in search matrix II %t\n", solutions.SearchMatrix2([][]int{
		{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
	}, 5))
	fmt.Printf("The start indices of p's anagrams in the string s are %v\n", solutions.FindAnagrams("cbaebabacd", "abc"))
	fmt.Printf("One of s1's permutations is the substring s2 %t\n", solutions.CheckInclusion("ab", "eidbaooo"))
	fmt.Printf("The final value after operations is %d\n", solutions.FinalValueAfterOperations([]string{"--X", "X++", "X++"}))
	fmt.Printf("The maximum achievable number is %d\n", solutions.TheMaximumAchievableX(4, 1))
	fmt.Printf("Constructing string from binary tree yields %v\n", solutions.Tree2str(&utils.TreeNode{
		Val: 1, Left: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 4}}, Right: &utils.TreeNode{Val: 3},
	}))

	codec := solutions.CodecConstructor()
	tinyUrl := codec.Encode("https://leetcode.com/problems/design-tinyurl")
	fmt.Printf("The encoded tiny url is %s\n", tinyUrl)
	fmt.Printf("The original url is %s\n", codec.Decode(tinyUrl))

	fmt.Printf("All elements in two binary search trees are %v\n", solutions.GetAllElements(
		&utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 4}},
		&utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 3}},
	))
	fmt.Printf("The number of triplets that can form two arrays of Equal XOR are %d\n", solutions.CountTriplets([]int{2, 3, 1, 6, 7}))
	fmt.Printf("The number of employees who met the target are %d\n", solutions.NumberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Printf("The simplified canonical path is %s\n", solutions.SimplifyPath("/home//foo/"))
	fmt.Printf("Subtracting the product and sum of digits of an integer yields %d\n", solutions.SubtractProductAndSum(234))
	// fmt.Printf("The goal parser's interpretation of command is %s\n", solutions.Interpret("G()(al)"))
	fmt.Printf("The Lowest common ancestor of the deepest leaves is %v\n", solutions.LcaDeepestLeaves(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 6}, Right: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 4}}},
		Right: &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 8}},
	}))
	fmt.Printf("The smallest tree with all deepest nodes is %v\n", solutions.SubtreeWithAllDeepest(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 6}, Right: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 4}}},
		Right: &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 8}},
	}))
	fmt.Printf("The number of contiguous subarrays where product of all elements is strictly less than k is %d\n", solutions.NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Printf("The subarray sums divisible by k are %d\n", solutions.SubarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	fmt.Printf("The amount of numbers smaller than the current number are %v\n", solutions.SmallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Printf("The Range sum of the Binary search tree is %d\n", solutions.RangeSumBST(&utils.TreeNode{
		Val: 10, Left: &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 7}},
		Right: &utils.TreeNode{Val: 15, Right: &utils.TreeNode{Val: 18}},
	}, 7, 15))
	fmt.Printf("The number of steps to reduce number to zero are %d\n", solutions.NumberOfSteps(14))
	fmt.Printf("The Element appearing more than 25 percent in sorted array is %d\n", solutions.FindSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	fmt.Printf("Sorting the list using insertion sort yields %v\n", solutions.InsertionSortList(&utils.ListNode{
		Val: 4, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 3}}},
	}))

	bSTIterator := solutions.BSTIteratorConstructor(&utils.TreeNode{
		Val:   7,
		Left:  &utils.TreeNode{Val: 3},
		Right: &utils.TreeNode{Val: 15, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 20}},
	})
	fmt.Printf("The Next value in BST is %d\n", bSTIterator.Next())
	fmt.Printf("The Next value in BST is %d\n", bSTIterator.Next())
	fmt.Printf("The BST has a next Value %t\n", bSTIterator.HasNext())
	fmt.Printf("The Next value in BST is %d\n", bSTIterator.Next())
	fmt.Printf("The BST has a next Value %t\n", bSTIterator.HasNext())
	fmt.Printf("The Next value in BST is %d\n", bSTIterator.Next())
	fmt.Printf("The BST has a next Value %t\n", bSTIterator.HasNext())
	fmt.Printf("The Next value in BST is %d\n", bSTIterator.Next())
	fmt.Printf("The BST has a next Value %t\n", bSTIterator.HasNext())

	fmt.Printf("The minimum number of moves to make array elements equal is %d\n", solutions.MinMoves([]int{1, 2, 3}))
	fmt.Printf("The minimum number of moves to make array elements equal II is %d\n", solutions.MinMoves2([]int{1, 2, 3}))
	fmt.Printf("The max product of two elements in an array is %d\n", solutions.MaxProduct3([]int{3, 4, 5, 2}))
	fmt.Printf("The max difference between node and ancestor is %d\n", solutions.MaxAncestorDiff(&utils.TreeNode{
		Val:   8,
		Left:  &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 4}, Right: &utils.TreeNode{Val: 7}}},
		Right: &utils.TreeNode{Val: 10, Right: &utils.TreeNode{Val: 14, Left: &utils.TreeNode{Val: 13}}},
	}))
	fmt.Printf("The minimum swaps to make the string balanced are %d\n", solutions.MinSwaps("][]["))
	fmt.Printf("The minimum add to make parentheses valid is %d\n", solutions.MinAddToMakeValid("())"))
	fmt.Printf("The string is an acronym of words %t\n", solutions.IsAcronym([]string{"alice", "bob", "charlie"}, "abc"))
	fmt.Printf("The number of special positions in the matrix are %d\n", solutions.NumSpecial([][]int{
		{1, 0, 0}, {0, 0, 1}, {1, 0, 0},
	}))
	fmt.Printf("The difference between ones and zeroes in a row and column is %v\n", solutions.OnesMinusZeros([][]int{
		{0, 1, 1}, {1, 0, 1}, {0, 0, 1},
	}))
	fmt.Printf("The minimum remove to make valid parentheses is %s\n", solutions.MinRemoveToMakeValid("))(("))
	fmt.Printf("The number of consistent strings are %d\n", solutions.CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Printf("The number of pairs of similar strings is %d\n", solutions.SimilarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
	fmt.Printf("The balanced binary search tree is %v\n", solutions.BalanceBST(&utils.TreeNode{
		Val: 1, Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3, Right: &utils.TreeNode{Val: 4}}},
	}))
	fmt.Printf("The result after separating the digits in the array is %v\n", solutions.SeparateDigits([]int{13, 25, 83, 77}))
	fmt.Printf("Removing trailing zeros from string yields %s\n", solutions.RemoveTrailingZeros("51230100"))
	fmt.Printf("The Kth smallest element inn a sorted matrix is %d\n", solutions.KthSmallest2([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 8))
	fmt.Printf("The greatest common divisor of the array is %d\n", solutions.FindGCD([]int{2, 5, 6, 9, 10}))
	fmt.Printf("The destination city is %s\n", solutions.DestCity([][]string{
		{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"},
	}))
	fmt.Printf("The ordering of the deck is %v\n", solutions.DeckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Printf("The minimum number of vertices to reach all nodes are %v\n", solutions.FindSmallestSetOfVertices(6, [][]int{
		{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
	}))
	fmt.Printf("The number of strings with prefix are %d\n", solutions.PrefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	fmt.Printf("The halves are alike %t\n", solutions.HalvesAreAlike("book"))
	fmt.Printf("The integer n is a sum of powers of three %t\n", solutions.CheckPowersOfThree(12))
	fmt.Printf("S is a good strings %t\n", solutions.AreOccurrencesEqual("abacbc"))

	foodRatings := solutions.FoodRatingsConstructor(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)

	fmt.Printf("The highest rated food is %s\n", foodRatings.HighestRated("korean"))
	fmt.Printf("The highest rated food is %s\n", foodRatings.HighestRated("japanese"))
	foodRatings.ChangeRating("sushi", 16)
	fmt.Printf("The highest rated food is %s\n", foodRatings.HighestRated("japanese"))
	foodRatings.ChangeRating("ramen", 16)
	fmt.Printf("The highest rated food is %s\n", foodRatings.HighestRated("japanese"))
	fmt.Printf("The numbers with even digits are %d\n", solutions.FindNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Printf("The maximum product difference between two pairs is %d\n", solutions.MaxProductDifference([]int{5, 6, 2, 7, 4}))
	fmt.Printf("The number of points inside a circle are %v\n", solutions.CountPoints(
		[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
	))
	fmt.Printf("The N-repeated element in size 2N Array is %d\n", solutions.RepeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Printf("The image after applying the smoother is %v\n", solutions.ImageSmoother([][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}))
	fmt.Printf("The maximum ice cream bars the boy can buy are %d\n", solutions.MaxIceCream([]int{1, 3, 2, 4, 1}, 7))
	fmt.Printf("The smallest even multiple of 2 and n is %d\n", solutions.SmallestEvenMultiple(5))
	fmt.Printf("The amount of money you will have after buying two chocolates is %d\n", solutions.BuyChoco([]int{1, 2, 2}, 3))

	nums3 := solutions.ShuffleConstructor([]int{1, 2, 3})
	fmt.Printf("Shuffle the array [1,2,3] and return its result %v\n", nums3.Shuffle())
	fmt.Printf("Resets the array back to its original configuration [1,2,3]. Return %v\n", nums3.Reset())
	fmt.Printf("Shuffle the array [1,2,3] and return its result %v\n", nums3.Shuffle())

	randomNodeSolution := solutions.RandomNodeConstructor(&utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 3}}})
	fmt.Printf("The random node is %d\n", randomNodeSolution.GetRandom())
	fmt.Printf("The random node is %d\n", randomNodeSolution.GetRandom())
	fmt.Printf("The random node is %d\n", randomNodeSolution.GetRandom())
	fmt.Printf("The random node is %d\n", randomNodeSolution.GetRandom())

	fmt.Printf("The widest vertical area between two points such that no points are inside the area is %d\n", solutions.MaxWidthOfVerticalArea([][]int{
		{8, 7}, {9, 9}, {7, 4}, {9, 7},
	}))
	fmt.Printf("The minimum number of operations to move all in each box is %v\n", solutions.MinOperations3("001011"))
	fmt.Printf("The first letter to appear twice is %v\n", solutions.RepeatedCharacter("abccbaacz"))
	fmt.Printf("The maximum score after splitting a string is %d\n", solutions.MaxScore("00"))
	fmt.Printf("Reversing substrings between each pair of parenthenses yields %s\n", solutions.ReverseParentheses("(ed(et(oc))el)"))
	fmt.Printf("The center of the graph is %d\n", solutions.FindCenter([][]int{
		{1, 2}, {2, 3}, {4, 2},
	}))
	fmt.Printf("I have visited this location before %t\n", solutions.IsPathCrossing("NESWW"))
	fmt.Printf("The minimum changes to make alternating binary string is %d\n", solutions.MinOperations2("0100"))
	fmt.Printf("The number of ways to decode string are %d\n", solutions.NumDecodings("226"))
	fmt.Printf("The number of possible ways to roll the dice, so the sum of numbers equals target is %d\n", solutions.NumRollsToTarget(2, 6, 7))
	fmt.Printf("The minimum time Bob needs to make the rope colorful is %d\n", solutions.MinCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Printf("The minimum length of the run-length encoded version of s after deleting atmost k characters %d\n", solutions.GetLengthOfOptimalCompression("aaabcccd", 2))
	fmt.Printf("The minimum difficulty of a job schedule is %d\n", solutions.MinDifficulty([]int{6, 5, 4, 3, 4, 2, 1}, 2))
	fmt.Printf("Every string can be made equal using any number of operartions %t\n", solutions.MakeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Printf("The largest substring between two equal characters is %d\n", solutions.MaxLengthBetweenEqualCharacters("abca"))
	fmt.Printf("The output after assigning cookies is %d\n", solutions.FindContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Printf("The total number of laser beams in the bank %d\n", solutions.NumberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	fmt.Printf("The target indices after sorting the array is %v\n", solutions.TargetIndices([]int{1, 2, 5, 2, 3}, 2))
	flatten := &utils.TreeNode{
		Val: 1, Left: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 4}}, Right: &utils.TreeNode{Val: 5, Right: &utils.TreeNode{Val: 6}},
	}
	solutions.Flatten(flatten)
	fmt.Printf("Flatten the binary tree to linked list yields %v\n", flatten)
	fmt.Printf("The maximum profit you can take is %d\n", solutions.JobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	fmt.Printf("The number of arithmetic slices II are %d\n", solutions.NumberOfArithmeticSlices2([]int{2, 4, 6, 8, 10}))
	fmt.Printf("The complement of base 10 integer is %d\n", solutions.BitwiseComplement(5))
	fmt.Printf("The complement of a number is %d\n", solutions.BitwiseComplement(1))
	fmt.Printf("The list of integers representing the size of the parts is %v\n", solutions.PartitionLabels("ababcbacadefegdehijhklij"))
	fmt.Printf("The number of minutes for the entire tree to be infected are %d\n", solutions.AmountOfTime(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 5, Right: &utils.TreeNode{Val: 4, Left: &utils.TreeNode{Val: 9}, Right: &utils.TreeNode{Val: 2}}},
		Right: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 10}, Right: &utils.TreeNode{Val: 6}},
	}, 3))
	fmt.Printf("The best time to buy and sell stock II is %d\n", solutions.MaxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("The percentage of letter in string is %d\n", solutions.PercentageLetter("foobar", 'o'))
	fmt.Printf("The players with zero or one losses are %v\n", solutions.FindWinners([][]int{
		{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9},
	}))

	randomisedSet := solutions.RandomizedSetConstructor()
	fmt.Printf("Inserts 1 to the set. Returns true as 1 was inserted successfully %t\n", randomisedSet.Insert(1))
	fmt.Printf("Returns false as 2 does not exist in the set. %t\n", randomisedSet.Remove(2))
	fmt.Printf("Inserts 2 to the set, returns true. Set now contains [1,2]. %t\n", randomisedSet.Insert(2))
	fmt.Printf("getRandom() should return either 1 or 2 randomly. %d\n", randomisedSet.GetRandom())
	fmt.Printf("Removes 1 from the set, returns true. Set now contains [2]. %t\n", randomisedSet.Remove(1))
	fmt.Printf("2 was already in the set, so return false. %t\n", randomisedSet.Insert(2))
	fmt.Printf("Since 2 is the only number in the set, getRandom() will always return 2. %d\n", randomisedSet.GetRandom())
	fmt.Printf("Replacing all digits with characters yields %s\n", solutions.ReplaceDigits("a1c1e1"))
	fmt.Printf("The minimum sum of any falling path is %d\n", solutions.MinFallingPathSum([][]int{
		{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
	}))
	fmt.Printf("The sum of subarray minimums is %v\n", solutions.SumSubarrayMins([]int{11, 81, 94, 43, 3}))
	fmt.Printf("The maximum amount of money you can rob without alerting the police is %d\n", solutions.Rob([]int{2, 7, 9, 3, 1}))
	fmt.Printf("The number that occurs twice and the number that is missing are %v\n", solutions.FindErrorNums([]int{1, 2, 2, 4}))
	fmt.Printf("The maximum possible length is %d\n", solutions.MaxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Printf("The number of pseudo-palindromic paths in a binary tree are %d\n", solutions.PseudoPalindromicPaths(&utils.TreeNode{
		Val:   2,
		Left:  &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 1}},
		Right: &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 1}},
	}))
	fmt.Printf("The length of two strings longest common subsequence is %d\n", solutions.LongestCommonSubsequence("abcde", "ace"))
	fmt.Printf("The number of paths are %d\n", solutions.FindPaths(2, 2, 2, 0, 0))
	fmt.Printf("The k inverse pairs of an array %d\n", solutions.KInversePairs(3, 0))
	fmt.Printf("The number of non-empty submatrices that sum to target are %d\n", solutions.NumSubmatrixSumTarget([][]int{
		{0, 1, 0}, {1, 1, 1}, {0, 1, 0},
	}, 0))

	myQueue := solutions.MyQueueConstructor()
	myQueue.Push(1)                                   // queue is: [1]
	myQueue.Push(2)                                   // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Printf("Peek returns %d\n", myQueue.Peek())   // return 1
	fmt.Printf("Pop returns %d\n", myQueue.Pop())     // return 1, queue is [2]
	fmt.Printf("Empty returns %t\n", myQueue.Empty()) // return false

	myStack := solutions.MyStackConstructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Printf("Top returns %d\n", myStack.Top())     // return 2
	fmt.Printf("Pop returns %d\n", myStack.Pop())     // return 2
	fmt.Printf("Empty returns %t\n", myStack.Empty()) // return false

	fmt.Printf("The 2D array containing all the arrays is %v\n", solutions.DivideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Printf("The sequential digits in the range are %v\n", solutions.SequentialDigits(100, 300))
	fmt.Printf("The largest sum of the given array after partitioning is %d\n", solutions.MaxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	fmt.Printf("The minimum window substring of s is %s\n", solutions.MinWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("The least number of perfect squares that sum to n are %d\n", solutions.NumSquares(12))
	fmt.Printf("The Largest divisible subset of distinctive positive integers is %v\n", solutions.LargestDivisibleSubset([]int{1, 2, 3}))
	fmt.Printf("The maximum number of queries collection using both robots is %d\n", solutions.CherryPickup([][]int{
		{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1},
	}))
	fmt.Printf("The largest possible perimeter of a polygon whose sides can be formed is %d\n", solutions.LargestPerimeter([]int{1, 12, 1, 2, 5, 50, 3}))
	fmt.Printf("The Least number of unique ints after k removals is %d\n", solutions.FindLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
	fmt.Printf("The farthest building index you can reach is %d\n", solutions.FurthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Printf("The number of rooms that held the most meetings are %d\n", solutions.MostBooked(2, [][]int{
		{0, 10}, {1, 5}, {2, 7}, {3, 4},
	}))
	fmt.Printf("The bitwise AND of all numbers in the range inclusive s %d\n", solutions.RangeBitwiseAnd(5, 7))
	fmt.Printf("The town judge is %d\n", solutions.FindJudge(2, [][]int{{1, 2}}))
	fmt.Printf("The cheapest flights within k stops are %d\n", solutions.FindCheapestPrice(4, [][]int{
		{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200},
	}, 0, 3, 1))
	fmt.Printf("The list of people that have a secret after all meetings have been taken are %v\n", solutions.FindAllPeople(6, [][]int{
		{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
	}, 1))
	fmt.Printf("It is possible to traverse between all such pairs of indices %t\n", solutions.CanTraverseAllPairs([]int{2, 3, 6}))
	fmt.Printf("The maximum possible score you can achieve is %d\n", solutions.BagOfTokensScore([]int{100, 200, 300, 400}, 200))
	fmt.Printf("The minimum length of a string after performing the operations is %d\n", solutions.MinimumLength("cabaabac"))
	fmt.Printf("The number of elements with maximum frequency are %d\n", solutions.MaxFrequencyElements([]int{1, 2, 2, 3, 1, 4}))
	fmt.Printf("The minimum integer common to both arrays is %d\n", solutions.GetCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))
	fmt.Printf("After repeatedly deleting consecutive sequences of nodes that sum to zero, the head of the list %v\n", solutions.RemoveZeroSumSublists(
		&utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: -3, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 1}}}}},
	))
	fmt.Printf("The pivot integer is %d\n", solutions.PivotInteger(8))
	fmt.Printf("The number of sub arrays with sum are %d\n", solutions.NumSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
	fmt.Printf("The minimum number of intervals required to complete all tasks is %d\n", solutions.LeastInterval([]byte{
		'A', 'A', 'A', 'B', 'B', 'B',
	}, 2))
	fmt.Printf("The first missing postive integer in an unsorted array is %d\n", solutions.FirstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Printf("The length of the longest subarray with at most k frequency is %d\n", solutions.MaxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
	fmt.Printf("The number of subarrays where max element appears atleast k times is %d\n", solutions.CountSubarrays([]int{1, 3, 2, 3, 3}, 2))
	fmt.Printf("The number of subarrays with fixed bounds is %d\n", solutions.CountSubarrays2([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	fmt.Printf("The number of subarrays with k different integers are %d\n", solutions.SubarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	fmt.Printf("The maximum depth of the parentheses is %d\n", solutions.MaxDepth2("(1+(2*3)+((8)/4))+1"))
	fmt.Printf("The final string after making it good is %s\n", solutions.MakeGood("leEeetcode"))
	fmt.Printf("The string is a valid parentheses string %t\n", solutions.CheckValidString("(*))"))
	fmt.Printf("The number of students unable to eat lunch are %d\n", solutions.CountStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	fmt.Printf("The time taken for a person at position k to finish buying tickets is %d\n", solutions.TimeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3))
	fmt.Printf("The smallest possible integer after removing k digits is %s\n", solutions.RemoveKdigits("1432219", 3))
	fmt.Printf("The amount of water an elevation can trap after raining is %d\n", solutions.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("The largest rectangle containing only 1s and 0s is %d\n", solutions.MaximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	fmt.Printf("Adding one row to the tree yields %v\n", solutions.AddOneRow(&utils.TreeNode{
		Val:   4,
		Left:  &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 1}},
		Right: &utils.TreeNode{Val: 6, Left: &utils.TreeNode{Val: 5}},
	}, 1, 2))
	fmt.Printf("The smallest string starting from leaf is %s\n", solutions.SmallestFromLeaf(&utils.TreeNode{
		Val:   0,
		Left:  &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 4}},
		Right: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 3}, Right: &utils.TreeNode{Val: 4}},
	}))
	fmt.Printf("The 2D array containing 4 length arrays is %v\n", solutions.FindFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Printf("There is a valid path from source to destination %t\n", solutions.ValidPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Printf("The minimum total turns required to open the lock are %d\n", solutions.OpenLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Printf("The list of all minimum height tree root labels is %v\n", solutions.FindMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Printf("The length of the longest Ideal string is %d\n", solutions.LongestIdealString("acfgbd", 2))
	fmt.Printf("The minimum sum of any falling path II is %d\n", solutions.MinFallingPathSum2([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
	fmt.Printf("The minimum numbers of steps to space all the characters in the keyword are %d\n", solutions.FindRotateSteps("godding", "gd"))
	fmt.Printf("The sum of the distances between the ith node in the tree and all other nodes is %v\n", solutions.SumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	fmt.Printf("The minimum number of operations required to make the bitwise XOR of all elements of the final array equal to k are %d\n", solutions.MinOperations4([]int{2, 1, 3, 4}, 1))
	fmt.Printf("The number of wonderful non-empty substrings in word are %d\n", solutions.WonderfulSubstrings("aabb"))
	fmt.Printf("The largest positive integer that exists with it's negative is %d\n", solutions.FindMaxK([]int{-1, 2, -3, 3}))
	fmt.Printf("The result after comparing two version numbers is %d\n", solutions.CompareVersion("1.01", "1.001"))
	fmt.Printf("The minimum number of boats to carry every given person is %d\n", solutions.NumRescueBoats([]int{1, 2}, 3))
	fmt.Printf("The head of a linked list after doubling it is %v\n", solutions.DoubleIt(&utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 8, Next: &utils.ListNode{Val: 9}}}))
	fmt.Printf("The array answer of size n where answer[i] is the rank of the ith athlete is %v\n", solutions.FindRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Printf("The maximum sum of the happiness values of the selected children you can achieve by selecting k children. is %d\n", solutions.MaximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Printf("The kth smallest fraction considered is %v\n", solutions.KthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Printf("The least amount of money needed to form a paid group of n people is %v\n", solutions.MincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
	fmt.Printf("The largest local values in a matrix are %v\n", solutions.LargestLocal([][]int{
		{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2},
	}))
	fmt.Printf("The highest possible score after making any number of moves %d\n", solutions.MatrixScore([][]int{
		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	}))
	fmt.Printf("The maximum amount of gold you can collect %d\n", solutions.GetMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
	fmt.Printf("The safest path in the grid is %d\n", solutions.MaximumSafenessFactor([][]int{
		{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0},
	}))
	fmt.Printf("The boolean result of evaluating the root node is %t\n", solutions.EvaluateTree(&utils.TreeNode{
		Val:   2,
		Left:  &utils.TreeNode{Val: 1},
		Right: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 1}},
	}))
	fmt.Printf("Removing all leaf nods with the target value yields %v\n", solutions.RemoveLeafNodes(&utils.TreeNode{
		Val:   1,
		Left:  &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 2}},
		Right: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 2}, Right: &utils.TreeNode{Val: 4}}},
	}, 2))
	fmt.Printf("The minimum number of moves required to make every node have exactly one coin is %d\n", solutions.DistributeCoins(&utils.TreeNode{
		Val: 3, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 0},
	}))
	fmt.Printf("The maximum possible sum of the values Alice can achieve by performing the operation any number of times is %d\n", solutions.MaximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
	fmt.Printf("The sum of all subset XOR totals is %d\n", solutions.SubsetXORSum([]int{1, 3}))
	fmt.Printf("The number of non-empty beautiful subsets of the array nums is %d\n", solutions.BeautifulSubsets([]int{2, 4, 6}, 2))
	fmt.Printf("The maximum score of any valid set of words formed are %d\n", solutions.MaxScoreWords(
		[]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
		[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
	fmt.Printf("The possible sentences in the word dictionary are %v\n", solutions.WordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Printf("The number of possible attendance records of length n are %d\n", solutions.CheckRecord(2))
	fmt.Printf("It can be proven that the array is special since there are %d numbers greater than or equal to the length of the array\n", solutions.SpecialArray([]int{3, 5}))
	fmt.Printf("The maximum length of a substring of s that can be changed is %d\n", solutions.EqualSubstring("abcd", "bcdf", 3))
	fmt.Printf("The number of steps to reduce a binary representation to one is %d\n", solutions.NumSteps("1101"))
	fmt.Printf("The score of the string is %d\n", solutions.ScoreOfString("hello"))
	fmt.Printf("The minimum number of characters that need to be appended %d\n", solutions.AppendCharacters("coaching", "coding"))
	fmt.Printf("The longest palindrome that cann be built is %d\n", solutions.LongestPalindrome2("abccccdd"))
	fmt.Printf("The array of characters that show up in all strings are %v\n", solutions.CommonChars([]string{"bella", "label", "roller"}))
	fmt.Printf("She can rearrange the cards into groups so that each group is of size groupSize %t\n", solutions.IsNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

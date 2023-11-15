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
}

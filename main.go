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

	athletes := methods.New()

	athletes.Push(methods.Person{Name: "Lebron", Age: 37})
	athletes.Push(methods.Person{Name: "Jordan", Age: 50})
	athletes.Push(methods.Person{Name: "Kareem", Age: 65})

	athletes.Pop()

	fmt.Printf("There are %d GOATS\n", athletes.Length())

	nums := []int{1, 2, 3, 1}
	fmt.Printf("Array contains duplicate: %v\n", leetcode.ContainsDuplicate(nums))
	fmt.Printf("Is Anagram: %v\n", leetcode.IsAnagram("anagram", "nagaram"))
	fmt.Printf("Two sum indices are %v\n", leetcode.TwoSum([]int{-3, 4, 3, 90}, 0))
	fmt.Printf("Group Anagrams: %v\n", leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("The top k frequent elements are: %v\n", leetcode.TopKFrequent([]int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}, 4))
	fmt.Printf("The product of array except self is: %v\n", leetcode.ProductExceptSelf([]int{1, 2, 3, 4}))
}

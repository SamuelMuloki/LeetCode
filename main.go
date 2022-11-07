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

	fmt.Printf("There are %d GOATS", athletes.Length())

	nums := []int{1, 2, 3, 1}
	fmt.Printf("Array contains duplicate: %v", leetcode.ContainsDuplicate(nums))
}

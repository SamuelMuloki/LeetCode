package main

import (
	"fmt"

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
}

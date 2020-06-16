package main

import (
	"fmt"

	"github.com/SamuelMuloki/GOExamples/methods"
	"github.com/SamuelMuloki/GOExamples/utils"
)

func main() {
	a, b := utils.Swap("hello", "world")

	animals := []methods.Animal{methods.Dog{}, methods.Cat{}, methods.Llama{}, methods.JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	fmt.Println("hello world")
	fmt.Println("Hello, 世界")
	fmt.Println(a, b)
}

package main

import (
	"fmt"

	"github.com/SamuelMuloki/GOExamples/utils"
)

func main() {
	a, b := utils.Swap("hello", "world")
	fmt.Println("hello world")
	fmt.Println("Hello, 世界")
	fmt.Println(a, b)
}

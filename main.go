package main

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

func Sum(a int, b int) int {
	return a + b
}

func main() {
	result := Multiply(1, 2)
	fmt.Println(result)

	resSum := Sum(1, 2)
	fmt.Println(resSum)

	fmt.Println("zrdarova!!!")
	fmt.Println("Hello World")
}

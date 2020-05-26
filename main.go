package main

import (
	"fmt"
	"strings"

	"github.com/treasureBear94/learngo/something"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("a", "b", "c", "d", "e")

	totalLength, upperName := lenAndUpper("treasure")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("treasure")
	fmt.Println(totalLength2)

	fmt.Println(multiply(2, 2))

	fmt.Println("Hello world!")
	something.SayHello()

	const name string = "bear"
	var name2 string = "bear"
	name3 := "bear"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
}

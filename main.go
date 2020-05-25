package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello()

	const name string = "bear"
	var name2 string = "bear"
	name3 := "bear"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
}

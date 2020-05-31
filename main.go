package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 5
	fmt.Println(&a, b)
	fmt.Println(*b)
	*b = 20
	fmt.Print(a)
}

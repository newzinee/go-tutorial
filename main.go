package main

import "fmt"

func main() {
	names := [5]string{"a", "b", "c"}
	names[3] = "d"
	fmt.Println(names)

	empty := [5]string{}
	empty[0] = "1"
	fmt.Println(empty)

	names2 := []string{"a", "b", "c"}
	names2 = append(names2, "d")

	fmt.Println(names2)
}

package main

import "fmt"

func main() {
	trea := map[string]string{"name": "trea", "age": "20"}
	fmt.Println(trea)

	for key, value := range trea {
		fmt.Println(key, value)
	}
}

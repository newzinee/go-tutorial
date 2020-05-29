package main

import (
	"fmt"
)

func supperAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func supperAdd0(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	return 1
}

func supperAdd(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	supperAdd2(1, 2, 3, 4, 5, 6)
	fmt.Println("----------")
	supperAdd0(1, 2, 3, 4, 5, 6)
	fmt.Println("----------")
	result := supperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

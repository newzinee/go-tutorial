package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "apple"}
	trea := person{"trea", 18, favFood}
	trea2 := person{name: "trea", age: 18, favFood: favFood}
	fmt.Println(trea.name)
	fmt.Println(trea2.name)
}

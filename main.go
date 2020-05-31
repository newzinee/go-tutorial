package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink2(age int) bool {
	switch {
	case age > 18:
		return true
	case age == 18:
		return true
	}
	return false
}

func canIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(16))
}

package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}

package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/accounts"
	"github.com/treasureBear94/learngo/mydict"
)

func main() {
	account := accounts.NewAccount("treasure")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}

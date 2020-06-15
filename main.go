package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/accounts"
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
	fmt.Println(account.Balance())
}

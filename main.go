package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("treasure")
	account.Deposit(10)
	fmt.Println(account.Balance())
}

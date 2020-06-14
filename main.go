package main

import (
	"fmt"

	"github.com/treasureBear94/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("treasure")
	fmt.Println(account)
}

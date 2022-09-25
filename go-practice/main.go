package main

import (
	"fmt"

	"github.com/roha-lee/go-study/accounts"
)

func main() {
	account := accounts.NewAccount("roha")
	account.ShowAccount()
	account.Deposit(1000)
	account.ShowAccount()
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	account.ShowAccount()
}

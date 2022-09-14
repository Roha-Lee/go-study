package main

import (
	"fmt"

	"github.com/roha-lee/go-study/accounts"
)

func main() {
	account := accounts.NewAccount("roha")
	fmt.Println(account)
}

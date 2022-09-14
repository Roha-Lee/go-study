package main

import (
	"fmt"
	"std/github.com/roha-lee/go-study/banking"
)

func main() {
	account := banking.Account{Owner: "Roha Lee", Balance: 10000000}
	fmt.Println(account)
}

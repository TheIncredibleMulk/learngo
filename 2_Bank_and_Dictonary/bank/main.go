package main

import (
	"fmt"

	"github.com/TheIncredibleMulk/learngo/2_Bank_and_Dictonary/accounts"
)

func main() {
	account := accounts.NewAccount("Andrew")
	account.Deposit(10)
	err := account.Withdraw(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

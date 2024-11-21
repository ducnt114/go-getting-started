package main

import (
	"fmt"
	"time"
)

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.Balance <= amount {
		return false
	}
	b.Balance -= amount
	return true
}

func main() {
	account := BankAccount{
		Balance: 100,
	}

	for i := 0; i < 50; i++ {
		go account.Deposit(2)
	}
	for i := 0; i < 100; i++ {
		go account.Withdraw(1)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(account.Balance)
	// TODO: ensure the final balance is 100
}

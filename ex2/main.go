package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	m       *sync.RWMutex
	m2      *sync.Mutex
	Balance int32
}

func (b *BankAccount) Deposit(amount int32) {
	//atomic.AddInt32(&b.Balance, amount)
	b.m.Lock()
	// b.m.RLock()
	b.Balance += amount
	b.m.Unlock()
}

func (b *BankAccount) Withdraw(amount int32) bool {
	if b.Balance <= amount {
		return false
	}
	b.Balance -= amount
	return true
}

func main() {
	account := BankAccount{
		Balance: 100,
		m:       &sync.RWMutex{},
		m2:      &sync.Mutex{},
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

package main

import (
	"fmt"
	"sync"
)

/*
Banking System Implementation

Prompt:
Implement a basic banking system with the following functionalities:

create_account(account_id): Initialize a new bank account.
deposit(account_id, amount): Add funds to an account.
withdraw(account_id, amount): Remove funds from an account, ensuring sufficient balance.
transfer(from_account_id, to_account_id, amount): Move funds between accounts.
get_balance(account_id): Retrieve the current balance of an account.

Considerations:
Concurrency control to handle simultaneous transactions
Data consistency and integrity
Error handling for invalid operations
*/

type Account struct {
	l sync.Mutex
	user    string
	balance float32 // we use float here to handle decimal values
}

type Bank struct {
	l sync.Mutex
	accounts map[string]*Account // [account_id]Account
}


func (b *Bank) create_account(account_id string) {
	b.l.Lock()
	defer b.l.Unlock()

	// create the account and add it to the accounts with the account_id
	acc := Account {
		user: account_id,
		balance: 0,
	}

	b.accounts[account_id] = &acc
}


func (b *Bank) deposit(account_id string, amount float32) {
	b.l.Lock()
	defer b.l.Unlock()

	v, ok := b.accounts[account_id]
		if !ok {
			fmt.Println("key not found")
	}

	newbalance := v.balance + amount
	v.balance = newbalance

	b.accounts[account_id] = v
}

func (b *Bank) withdraw(account_id string, amount float32) {
	b.l.Lock()
	defer b.l.Unlock()

	v, ok := b.accounts[account_id]
	if !ok {
		fmt.Println("key not found during withdrawl")
	}

	newBalance := v.balance - amount
	v.balance = newBalance

	b.accounts[account_id] = v
}


func (b *Bank) transfer(from_account_id string, to_account_id string, amount float32) {
	b.l.Lock()
	defer b.l.Unlock()

	_, okTo := b.accounts[to_account_id]
	if !okTo {
		fmt.Println("to account not found")
		return // return to exit
	}

	_, okFrom := b.accounts[from_account_id]
	if !okFrom {
		fmt.Println("from account is not found")
		return // return to exit
	}

	b.withdraw(from_account_id, amount)
	b.deposit(to_account_id, amount)
}

func (b *Bank) get_balance(account_id string) float32 {
	b.l.Lock()
	defer b.l.Unlock()
	v, ok := b.accounts[account_id]
	if !ok {
		fmt.Println("key not found when getting balance")
	}
	return v.balance
}

func main() {
	fmt.Println("hello, world")
}

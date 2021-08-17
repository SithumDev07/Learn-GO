package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Bank Balance is insufficient")

//NewAccount Creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposti x Amount in account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

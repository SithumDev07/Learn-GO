package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

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

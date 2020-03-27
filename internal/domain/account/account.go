package domain

type Customer struct {
	Name    string
	Address string
}

type Account struct {
	Customer
	Balance float64
	Id      int
}

func (a *Account) SetBalance(newBalance float64) {
	a.Balance = newBalance
}

func (a *Account) SetCustomer(name, address string) {
	a.Name = name
	a.Address = address
}

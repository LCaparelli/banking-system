package domain

type Repository interface {
	AccountById(id int) (*Account, error)
	DeleteAccountById(id int) error
	NewAccount(account *Account) error
	SaveBalance(account *Account, newBalance float64) error
}

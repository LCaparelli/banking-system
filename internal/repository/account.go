package repository

import (
	"fmt"
	"github.com/LCaparelli/banking-system/internal/domain/account"
)

const (
	accountNotFound = "account %d does not exist"
)

type AccountRepository struct {
	accounts map[int]*domain.Account
}

func InMemRepoFactory() *AccountRepository {
	return &AccountRepository{accounts: make(map[int]*domain.Account)}
}

func (r *AccountRepository) AccountById(id int) (*domain.Account, error) {
	account, ok := r.accounts[id]
	if !ok {
		return nil, fmt.Errorf(fmt.Sprintf(accountNotFound, id))
	}
	return account, nil
}

func (r *AccountRepository) DeleteAccountById(id int) error {
	_, ok := r.accounts[id]
	if !ok {
		return fmt.Errorf(fmt.Sprintf(accountNotFound, id))
	}
	delete(r.accounts, id)
	return nil
}

func (r *AccountRepository) NewAccount(account *domain.Account) error {
	r.accounts[account.Id] = account
	return nil
}

func (r *AccountRepository) SaveBalance(account *domain.Account, newBalance float64) error {
	ac, ok := r.accounts[account.Id]
	if !ok {
		return fmt.Errorf(fmt.Sprintf(accountNotFound, account.Id))
	}
	ac.Balance = newBalance
	return nil
}

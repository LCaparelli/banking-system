package service

import (
	"fmt"
	"github.com/LCaparelli/banking-system/internal/domain/account"
	"github.com/LCaparelli/banking-system/internal/repository"
)

var currentId int

type AccountService struct {
	repo domain.Repository
}

func AccountServiceFactory() *AccountService {
	return &AccountService{repo: repository.InMemRepoFactory()}
}

func (s *AccountService) GetAccount(id int) (*domain.Account, error) {
	account, err := s.repo.AccountById(id)
	if err != nil {
		return nil, fmt.Errorf("AccountById: %v", err)
	}
	return account, nil
}

func (s *AccountService) DeleteAccount(id int) error {
	if err := s.repo.DeleteAccountById(id); err != nil {
		return fmt.Errorf("DeleteAccountById: %v", err)
	}
	return nil
}

func (s *AccountService) CreateAccount(name, address string, balance float64) (*domain.Account, error) {
	account := &domain.Account{Customer: domain.Customer{Name: name, Address: address}, Balance: balance, Id: currentId}
	if err := s.repo.NewAccount(account); err != nil {
		return nil, fmt.Errorf("NewAccount: %v", err)
	}
	currentId++
	return account, nil
}

func (s *AccountService) setBalance(account *domain.Account, balance float64) error {
	err := s.repo.SaveBalance(account, balance)
	if err != nil {
		return fmt.Errorf("SaveBalance: %v", err)
	}
	return nil
}

func (s *AccountService) Withdraw(id int, amount float64) (float64, error) {
	account, err := s.repo.AccountById(id)
	if err != nil {
		return -1, fmt.Errorf("AccountById: %v", err)
	}
	newBalance := account.Balance - amount
	if newBalance < 0.0 {
		return account.Balance, fmt.Errorf("insuffient funds to withdraw %.2f", amount)
	}
	err = s.setBalance(account, newBalance)
	if err != nil {
		return account.Balance, fmt.Errorf("setBalance: %v", err)
	}
	return newBalance, nil
}

func (s *AccountService) Deposit(id int, amount float64) (float64, error) {
	account, err := s.repo.AccountById(id)
	if err != nil {
		return -1, fmt.Errorf("AccountById: %v", err)
	}

	newBalance := account.Balance + amount
	err = s.setBalance(account, newBalance)
	if err != nil {
		return account.Balance, fmt.Errorf("setBalance: %v", err)
	}
	return newBalance, nil
}

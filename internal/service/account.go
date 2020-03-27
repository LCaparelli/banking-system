package service

import (
	"banking-system/internal/domain/account"
	"banking-system/internal/repository"
	"fmt"
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

func (s *AccountService) CreateAccount(name, address string, balance float64) (int, error) {
	account := &domain.Account{Customer: domain.Customer{Name: name, Address: address}, Balance: balance, Id: currentId}
	if err := s.repo.NewAccount(account); err != nil {
		return -1, fmt.Errorf("NewAccount: %v", err)
	}
	currentId++
	return account.Id, nil
}

func (s *AccountService) setBalance(account *domain.Account, balance float64) error {
	err := s.repo.SaveBalance(account, balance)
	if err != nil {
		return fmt.Errorf("SaveBalance: %v", err)
	}
	return nil
}

func (s *AccountService) Withdraw(id int, amount float64) error {
	account, err := s.repo.AccountById(id)
	if err != nil {
		return fmt.Errorf("AccountById: %v", err)
	}
	newBalance := account.Balance - amount
	if newBalance < 0.0 {
		return fmt.Errorf("insuffient funds to withdraw %.2f", amount)
	}
	err = s.setBalance(account, newBalance)
	if err != nil {
		return fmt.Errorf("setBalance: %v", err)
	}
	return nil
}

func (s *AccountService) Deposit(id int, amount float64) error {
	account, err := s.repo.AccountById(id)
	if err != nil {
		return fmt.Errorf("AccountById: %v", err)
	}
	err = s.setBalance(account, account.Balance+amount)
	if err != nil {
		return fmt.Errorf("setBalance: %v", err)
	}
	return nil
}

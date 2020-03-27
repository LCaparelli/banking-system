package pkg

import (
	"fmt"
	"github.com/LCaparelli/banking-system/internal/domain/account"
	"github.com/LCaparelli/banking-system/internal/service"
)

var accountService = service.AccountServiceFactory()

func Account(id int) (*domain.Account, error) {
	account, err := accountService.GetAccount(id)
	if err != nil {
		return nil, fmt.Errorf("service.getaccount: %v", err)
	}
	return account, nil
}

func DeleteAccount(id int) error {
	if err := accountService.DeleteAccount(id); err != nil {
		return fmt.Errorf("service.deleteaccount: %v", err)
	}
	return nil
}

func CreateAccount(name, address string, balance float64) (int, error) {
	id, err := accountService.CreateAccount(name, address, balance)
	if err != nil {
		return -1, fmt.Errorf("service.createaccount: %v", err)
	}
	return id, nil
}

func Deposit(id int, amount float64) error {
	if err := accountService.Deposit(id, amount); err != nil {
		return fmt.Errorf("service.deposit: %v", err)
	}
	return nil
}

func Withdraw(id int, amount float64) error {
	if err := accountService.Withdraw(id, amount); err != nil {
		return fmt.Errorf("service.withdraw: %v", err)
	}
	return nil
}

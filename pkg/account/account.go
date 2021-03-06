package account

import (
	"fmt"
	"github.com/LCaparelli/banking-system/internal/domain/account"
	"github.com/LCaparelli/banking-system/internal/service"
	"github.com/LCaparelli/banking-system/internal/web/request"
)

var accountService = service.AccountServiceFactory()

func validateGet(id int) error {
	return request.AccountGET{Id: id}.Validate()
}

func validateDelete(id int) error {
	return request.AccountDELETE{Id: id}.Validate()
}

func validateCreate(name, address string, balance float64) error {
	return request.AccountPOST{Name: name, Address: address, Balance: balance}.Validate()
}

func validateDeposit(id int, amount float64) error {
	return request.DepositPOST{Id: id, Amount: amount}.Validate()
}

func validateWithdraw(id int, amount float64) error {
	return request.WithdrawPOST{Id: id, Amount: amount}.Validate()
}

func Account(id int) (*domain.Account, error) {
	if err := validateGet(id); err != nil {
		return nil, fmt.Errorf("validateget: %v", err)
	}
	account, err := accountService.GetAccount(id)
	if err != nil {
		return nil, fmt.Errorf("service.getaccount: %v", err)
	}
	return account, nil
}

func DeleteAccount(id int) error {
	if err := validateDelete(id); err != nil {
		return fmt.Errorf("validatedelete: %v", err)
	}
	if err := accountService.DeleteAccount(id); err != nil {
		return fmt.Errorf("service.deleteaccount: %v", err)
	}
	return nil
}

func CreateAccount(name, address string, balance float64) (*domain.Account, error) {
	if err := validateCreate(name, address, balance); err != nil {
		return nil, fmt.Errorf("validatecreate: %v", err)
	}
	account, err := accountService.CreateAccount(name, address, balance)
	if err != nil {
		return nil, fmt.Errorf("service.createaccount: %v", err)
	}
	return account, nil
}

func Deposit(id int, amount float64) (float64, error) {
	if err := validateDeposit(id, amount); err != nil {
		return -1, fmt.Errorf("validatedeposit: %v", err)
	}
	newBalance, err := accountService.Deposit(id, amount)
	if err != nil {
		return newBalance, fmt.Errorf("service.deposit: %v", err)
	}
	return newBalance, nil
}

func Withdraw(id int, amount float64) (float64, error) {
	if err := validateWithdraw(id, amount); err != nil {
		return -1, fmt.Errorf("validatewithdraw: %v", err)
	}
	newBalance, err := accountService.Withdraw(id, amount)
	if err != nil {
		return newBalance, fmt.Errorf("service.withdraw: %v", err)
	}
	return newBalance, nil
}

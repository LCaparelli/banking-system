package request

import "fmt"

type AccountGET struct {
	Id int
}

type AccountDELETE struct {
	Id int
}

type AccountPOST struct {
	Name    string
	Address string
	Balance float64
}

func (g *AccountGET) Validate() error {
	if g.Id < 0 {
		return fmt.Errorf(negativeId)
	}
	return nil
}

func (d *AccountDELETE) Validate() error {
	if d.Id < 0 {
		return fmt.Errorf(negativeId)
	}
	return nil
}

func (p *AccountPOST) Validate() error {
	if p.Name == "" || p.Address == "" {
		return fmt.Errorf(emptyNameOrAddress)
	}
	if p.Balance < 0 {
		return fmt.Errorf(negativeBalance)
	}
	return nil
}

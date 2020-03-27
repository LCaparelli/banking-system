package request

import "fmt"

type WithdrawPOST struct {
	Id     int
	Amount float64
}

func (p *WithdrawPOST) Validate() error {
	if p.Id < 0 {
		return fmt.Errorf(negativeId)
	}
	if p.Amount <= 0.0 {
		return fmt.Errorf(nonPositiveWithdraw)
	}
	return nil
}

package request

import "fmt"

type DepositPOST struct {
	Id     int
	Amount float64
}

func (p *DepositPOST) Validate() error {
	if p.Id < 0 {
		return fmt.Errorf(negativeId)
	}
	if p.Amount <= 0.0 {
		return fmt.Errorf(nonPositiveDeposit)
	}
	return nil
}

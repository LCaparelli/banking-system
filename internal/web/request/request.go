package request

const (
	emptyNameOrAddress  = "name and/or address is empty"
	negativeId          = "id is less than 0"
	negativeBalance     = "balance is less than 0.0"
	nonPositiveDeposit  = "the amount to deposit must be positive"
	nonPositiveWithdraw = "the amount to withdraw must be positive"
)

type Request interface {
	Validate() error
}

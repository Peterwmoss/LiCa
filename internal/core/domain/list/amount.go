package list

import "errors"

type Amount int

var (
	ErrInvalidAmount = errors.New("invalid amount")
)

func NewAmount(amount Amount) (Amount, error) {
	if amount < 1 {
		return 0, ErrInvalidAmount
	}

	return Amount(amount), nil
}

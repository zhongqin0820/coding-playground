package simple

import (
	"errors"
	"strconv"
)

// payment interface
type Payment interface {
	Pay(amount float32) string
}

// condition definition
const (
	Card = 1
	Cash = 2
)

// condition function
func GetPayment(m int) (Payment, error) {
	switch m {
	case Card:
		return &CardPay{}, nil
	case Cash:
		return &CashPay{}, nil
	default:
		return nil, errors.New("Nil implements")
	}
}

// concrete payment implementation
type CardPay struct{}

func (c *CardPay) Pay(amount float32) string {
	return strconv.FormatFloat(float64(amount), 'f', 2, 64)
}

type CashPay struct{}

func (p *CashPay) Pay(amount float32) string {
	return strconv.FormatFloat(float64(amount), 'f', 2, 64)
}
